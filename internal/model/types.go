package model

// ReceivedPost is the persisted and API-returned message shape after the server
// has normalized the incoming frontend payload.
type ReceivedPost struct {
	URL       string            `json:"url,omitempty" bson:"url,omitempty"`
	Content   string            `json:"content,omitempty" bson:"content,omitempty"`
	Username  string            `json:"username,omitempty" bson:"username,omitempty"`
	AvatarURL string            `json:"avatarUrl" bson:"avatar_url,omitempty"`
	TTS       bool              `json:"tts,omitempty" bson:"tts,omitempty"`
	TargetID  string            `json:"targetId,omitempty" bson:"targetId,omitempty"`
	Extras    map[string]string `json:"extras,omitempty" bson:"extras,omitempty"`
}

// IncomingPost matches the real frontend request body posted by index.js.
type IncomingPost struct {
	Content   string `json:"content"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	TTS       bool   `json:"tts"`
	Color     string `json:"color"`
	ImageURL  string `json:"imgUrl"`
	IP        string `json:"ip"`
	Thumbnail string `json:"thumbnail"`
	TargetID  string `json:"targetId"`
}

// StoreData is the history record stored in MongoDB and returned by GET /HideBot/discord.
type StoreData struct {
	ID             string       `json:"id,omitempty" bson:"_id,omitempty"`
	TimeStamp      int64        `json:"timeStamp" bson:"timeStamp"`
	SerialNumber   int          `json:"serialNumber" bson:"serialNumber"`
	DiscordMessage ReceivedPost `json:"discordMessage" bson:"discordMessage"`
	PosterIP       string       `json:"posterIp" bson:"posterIp"`
	ResponseCode   *string      `json:"responseCode" bson:"responseCode,omitempty"`
}

// SerialCounter stores the current anonymous post serial number.
type SerialCounter struct {
	ID      any `bson:"_id,omitempty"`
	Counter int `bson:"counter"`
}

// EmbedField represents one Discord embed field entry.
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

// EmbedAuthor represents the author block in a Discord embed.
type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// EmbedThumbnail represents the small thumbnail image shown in the embed.
type EmbedThumbnail struct {
	URL string `json:"url,omitempty"`
}

// EmbedImage represents the large content image shown in the embed.
type EmbedImage struct {
	URL string `json:"url,omitempty"`
}

// DiscordEmbed is the payload block forwarded to the Discord webhook API.
type DiscordEmbed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Color       int             `json:"color,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
}

// DiscordWebhookPayload is the outbound webhook body sent to Discord.
type DiscordWebhookPayload struct {
	URL       string         `json:"-"`
	Content   string         `json:"content,omitempty"`
	Username  string         `json:"username,omitempty"`
	AvatarURL string         `json:"avatar_url,omitempty"`
	TTS       bool           `json:"tts,omitempty"`
	Embeds    []DiscordEmbed `json:"embeds,omitempty"`
}

// HealthResponse mirrors the minimal /actuator/health response expected by Fly.io.
type HealthResponse struct {
	Status string `json:"status"`
}
