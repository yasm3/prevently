package domain

type DeviceType string

const (
	Discord DeviceType = "discord"
)

type DiscordConfig struct {
	WebhookURL string `json:"webhook_url"`
}

type Device struct {
	ID     string
	UserID string
	Name   string
	Type   DeviceType
	Config []byte
}
