package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"github.com/DF-wu/HideReplier/internal/store"
)

const (
	authorIconURL = "https://img.icons8.com/color/144/000000/drupal.png"
	colorKey      = "color"
	avatarURLKey  = "avatar_url"
	ipKey         = "ip"
	thumbnailKey  = "thumbnail"
	imageURLKey   = "imgUrl"
)

type DiscordService struct {
	config  config.Config
	store   *store.MongoStore
	client  *http.Client
	mu      sync.Mutex
	counter *model.SerialCounter
}

// NewDiscordService loads the persisted counter state and prepares the webhook client.
func NewDiscordService(ctx context.Context, cfg config.Config, mongoStore *store.MongoStore) (*DiscordService, error) {
	counter, err := mongoStore.LoadOrCreateCounter(ctx)
	if err != nil {
		return nil, err
	}

	return &DiscordService{
		config:  cfg,
		store:   mongoStore,
		client:  &http.Client{Timeout: 15 * time.Second},
		counter: counter,
	}, nil
}

// GetVersion returns the configured bot version string exposed to the frontend.
func (s *DiscordService) GetVersion() string {
	return s.config.BotVersion
}

// GetHistory returns the stored anonymous message history sorted by serial number.
func (s *DiscordService) GetHistory(ctx context.Context) ([]model.StoreData, error) {
	return s.store.ListHistory(ctx)
}

// PostAnonymousMessage normalizes the frontend payload, sends the webhook,
// and persists the history only after a successful dispatch.
func (s *DiscordService) PostAnonymousMessage(ctx context.Context, post model.IncomingPost) (model.ReceivedPost, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	normalized := model.ReceivedPost{
		Content:   post.Content,
		Username:  post.Username,
		AvatarURL: post.AvatarURL,
		TTS:       post.TTS,
		Extras: map[string]string{
			colorKey:     post.Color,
			avatarURLKey: post.AvatarURL,
			imageURLKey:  post.ImageURL,
			ipKey:        post.IP,
			thumbnailKey: post.Thumbnail,
		},
	}

	originalContent := normalized.Content
	posterIP := normalized.Extras[ipKey]
	normalizedColor, err := normalizeColor(normalized.Extras[colorKey])
	if err != nil {
		return model.ReceivedPost{}, err
	}

	normalized.AvatarURL = normalized.Extras[avatarURLKey]
	normalized.Extras[colorKey] = strconv.Itoa(normalizedColor)

	s.counter.Counter++
	serialNumber := s.counter.Counter
	timestamp := currentTaiwanEpochSecond()

	payload := model.DiscordWebhookPayload{
		URL:       s.config.DiscordWebhook,
		Content:   "",
		Username:  normalized.Username,
		AvatarURL: normalized.AvatarURL,
		TTS:       normalized.TTS,
		Embeds: []model.DiscordEmbed{
			buildEmbed(s.config, normalized, originalContent, normalizedColor, posterIP, serialNumber),
		},
	}

	if err := s.dispatch(ctx, payload); err != nil {
		s.counter.Counter--
		return model.ReceivedPost{}, err
	}

	normalized.URL = s.config.DiscordWebhook
	normalized.Content = originalContent

	storeData := model.StoreData{
		TimeStamp:      timestamp,
		SerialNumber:   serialNumber,
		DiscordMessage: normalized,
		PosterIP:       posterIP,
	}

	if err := s.store.InsertHistory(ctx, storeData); err != nil {
		s.counter.Counter--
		return model.ReceivedPost{}, err
	}

	if err := s.store.SaveCounter(ctx, s.counter); err != nil {
		s.counter.Counter--
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
	_, _ = io.ReadAll(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("discord webhook request failed with status %d", resp.StatusCode)
	}

	return nil
}

func buildEmbed(cfg config.Config, post model.ReceivedPost, content string, normalizedColor int, posterIP string, serialNumber int) model.DiscordEmbed {
	return model.DiscordEmbed{
		Title:       post.Username,
		Description: content,
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
			{Name: "來自：", Value: posterIP, Inline: true},
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

func currentTaiwanEpochSecond() int64 {
	location := time.FixedZone("Asia/Taipei", 8*60*60)
	return time.Now().In(location).Unix()
}
