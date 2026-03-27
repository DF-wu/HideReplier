package model

type ReceivedPost struct {
	URL       string            `json:"url,omitempty" bson:"url,omitempty"`
	Content   string            `json:"content,omitempty" bson:"content,omitempty"`
	Username  string            `json:"username,omitempty" bson:"username,omitempty"`
	AvatarURL string            `json:"avatarUrl" bson:"avatar_url,omitempty"`
	TTS       bool              `json:"tts,omitempty" bson:"tts,omitempty"`
	Extras    map[string]string `json:"extras,omitempty" bson:"extras,omitempty"`
}

type IncomingPost struct {
	Content   string `json:"content"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	TTS       bool   `json:"tts"`
	Color     string `json:"color"`
	ImageURL  string `json:"imgUrl"`
	IP        string `json:"ip"`
	Thumbnail string `json:"thumbnail"`
}

type StoreData struct {
	ID             string       `json:"id,omitempty" bson:"_id,omitempty"`
	TimeStamp      int64        `json:"timeStamp" bson:"timeStamp"`
	SerialNumber   int          `json:"serialNumber" bson:"serialNumber"`
	DiscordMessage ReceivedPost `json:"discordMessage" bson:"discordMessage"`
	PosterIP       string       `json:"posterIp" bson:"posterIp"`
	ResponseCode   *string      `json:"responseCode" bson:"responseCode,omitempty"`
}

type SerialCounter struct {
	ID      any `bson:"_id,omitempty"`
	Counter int `bson:"counter"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedThumbnail struct {
	URL string `json:"url,omitempty"`
}

type EmbedImage struct {
	URL string `json:"url,omitempty"`
}

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

type DiscordWebhookPayload struct {
	URL       string         `json:"-"`
	Content   string         `json:"content,omitempty"`
	Username  string         `json:"username,omitempty"`
	AvatarURL string         `json:"avatar_url,omitempty"`
	TTS       bool           `json:"tts,omitempty"`
	Embeds    []DiscordEmbed `json:"embeds,omitempty"`
}

type HealthResponse struct {
	Status string `json:"status"`
}
