package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
)

const (
	authorIconURL = "https://img.icons8.com/color/144/000000/drupal.png"
	colorKey      = "color"
	avatarURLKey  = "avatar_url"
	ipKey         = "ip"
	thumbnailKey  = "thumbnail"
	imageURLKey   = "imgUrl"

	// Discord never returns a meaningful body for webhook posts; cap how much
	// we are willing to drain so a misbehaving upstream cannot eat memory.
	maxWebhookResponseBytes = 64 << 10
)

var taiwanZone = time.FixedZone("Asia/Taipei", 8*60*60)

type DiscordService struct {
	config config.Config
	store  Store
	client *http.Client
}

type Store interface {
	NextSerial(ctx context.Context) (int, error)
	InsertHistory(ctx context.Context, data model.StoreData) error
	StreamHistory(ctx context.Context, limit int64, fn func(model.StoreData) error) error
}

// NewDiscordService prepares the webhook client. The HTTP transport keeps a
// couple of warm connections to Discord so each post does not pay a fresh
// TLS handshake on a CPU-starved machine.
func NewDiscordService(cfg config.Config, dataStore Store) *DiscordService {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        4,
		MaxIdleConnsPerHost: 2,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &DiscordService{
		config: cfg,
		store:  dataStore,
		client: &http.Client{Timeout: 15 * time.Second, Transport: transport},
	}
}

// GetVersion returns the configured bot version string exposed to the frontend.
func (s *DiscordService) GetVersion() string {
	return s.config.BotVersion
}

// StreamHistory streams stored anonymous messages sorted by serial number.
func (s *DiscordService) StreamHistory(ctx context.Context, limit int64, fn func(model.StoreData) error) error {
	return s.store.StreamHistory(ctx, limit, fn)
}

// PostAnonymousMessage normalizes the frontend payload, reserves a serial
// number, sends the webhook, and persists the history only after a successful
// dispatch. Requests are fully concurrent: the serial number is reserved
// atomically in the store, so a failed dispatch leaves a gap instead of
// serializing every post behind a global lock.
func (s *DiscordService) PostAnonymousMessage(ctx context.Context, post model.IncomingPost) (model.ReceivedPost, error) {
	normalizedColor, err := normalizeColor(post.Color)
	if err != nil {
		return model.ReceivedPost{}, err
	}

	target, err := s.config.ResolveDiscordTarget(post.TargetID)
	if err != nil {
		return model.ReceivedPost{}, err
	}

	normalized := model.ReceivedPost{
		Content:   post.Content,
		Username:  post.Username,
		AvatarURL: post.AvatarURL,
		TTS:       post.TTS,
		TargetID:  target.ID,
		Extras: map[string]string{
			colorKey:     strconv.Itoa(normalizedColor),
			avatarURLKey: post.AvatarURL,
			imageURLKey:  post.ImageURL,
			ipKey:        post.IP,
			thumbnailKey: post.Thumbnail,
		},
	}

	serialNumber, err := s.store.NextSerial(ctx)
	if err != nil {
		return model.ReceivedPost{}, fmt.Errorf("reserve serial number: %w", err)
	}
	timestamp := time.Now().In(taiwanZone).Unix()

	payload := model.DiscordWebhookPayload{
		URL:       target.WebhookURL,
		Username:  normalized.Username,
		AvatarURL: normalized.AvatarURL,
		TTS:       normalized.TTS,
		Embeds: []model.DiscordEmbed{
			buildEmbed(s.config, normalized, normalizedColor, serialNumber),
		},
	}

	if err := s.dispatch(ctx, payload); err != nil {
		return model.ReceivedPost{}, err
	}

	normalized.URL = s.config.HostURL

	storeData := model.StoreData{
		TimeStamp:      timestamp,
		SerialNumber:   serialNumber,
		DiscordMessage: normalized,
		PosterIP:       post.IP,
	}

	if err := s.store.InsertHistory(ctx, storeData); err != nil {
		return model.ReceivedPost{}, err
	}

	return normalized, nil
}

func (s *DiscordService) dispatch(ctx context.Context, payload model.DiscordWebhookPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, payload.URL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "HideReplier-Go/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Drain (bounded) so the connection can be reused by the transport.
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, maxWebhookResponseBytes))

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("discord webhook request failed with status %d", resp.StatusCode)
	}

	return nil
}

func buildEmbed(cfg config.Config, post model.ReceivedPost, normalizedColor int, serialNumber int) model.DiscordEmbed {
	return model.DiscordEmbed{
		Title:       post.Username,
		Description: post.Content,
		URL:         cfg.HostURL,
		Color:       normalizedColor,
		Author: &model.EmbedAuthor{
			Name:    "匿名機器人v" + cfg.BotVersion + "（點我去發文）",
			URL:     cfg.HostURL,
			IconURL: authorIconURL,
		},
		Thumbnail: &model.EmbedThumbnail{URL: post.Extras[thumbnailKey]},
		Image:     &model.EmbedImage{URL: post.Extras[imageURLKey]},
		Fields: []model.EmbedField{
			{Name: "流水號", Value: strconv.Itoa(serialNumber), Inline: true},
			{Name: "來自：", Value: post.Extras[ipKey], Inline: true},
		},
	}
}

func normalizeColor(value string) (int, error) {
	trimmed := strings.TrimPrefix(value, "#")
	if trimmed == "" {
		return 0, fmt.Errorf("color is required")
	}

	parsed, err := strconv.ParseInt(trimmed, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid color: %w", err)
	}

	return int(parsed), nil
}
