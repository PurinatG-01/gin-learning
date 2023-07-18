package service

import (
	"encoding/json"
	"fmt"
	model "gin-learning/models"
	"os"
	"time"

	sender "github.com/bensch777/discord-webhook-golang"
)

type DiscordService interface {
	SendTransactionMessage(transactionId string, amount int, status string) error
	SendEmbedMessage(title string, color int, fields []sender.Field, footer sender.Footer) error
	CreateEmbed(title string, color int, fields []sender.Field, footer sender.Footer) *sender.Embed
	SendMessage(embeds sender.Embed) error
}

func NewDiscordService() DiscordService {
	url := fmt.Sprintf(os.Getenv("DISCORD_HOOK_URL"))
	if url == "" {
		panic("DISCORD_HOOK_URL is not set")
	}
	app_env := fmt.Sprintf(os.Getenv("APP_ENV"))
	if app_env == "" {
		panic("APP_ENV is not set")
	}
	return &discordService{HookUrl: url, SenderName: fmt.Sprintf("[%s-logger]", app_env)}
}

type discordService struct {
	HookUrl    string
	SenderName string
}

func (s *discordService) MapTransactionStatusColor(status string) int {
	switch status {
	case string(model.OMISE_CHARGE_STATUS_PENDING):
		return 16776960
	case string(model.OMISE_CHARGE_STATUS_SUCCESSFUL):
		return 3066993
	case string(model.OMISE_CHARGE_STATUS_FAILED):
		return 15158332
	case string(model.OMISE_CHARGE_STATUS_REVERSED):
		return 15158332
	case model.OMISE_CHARGE_STATUS_EXPIRED:
		return 15158332
	default:
		return 0
	}
}

func (s *discordService) MapTransactionStatusIcon(status string) string {
	switch status {
	case string(model.OMISE_CHARGE_STATUS_PENDING):
		return "⏳"
	case string(model.OMISE_CHARGE_STATUS_SUCCESSFUL):
		return "✅"
	case string(model.OMISE_CHARGE_STATUS_FAILED):
		return "❌"
	case string(model.OMISE_CHARGE_STATUS_REVERSED):
		return "🔁"
	case model.OMISE_CHARGE_STATUS_EXPIRED:
		return "⏰"
	default:
		return ""
	}
}

func (s *discordService) SendTransactionMessage(transactionId string, amount int, status string) error {
	fields := []sender.Field{
		{
			Name:   "Transaction ID",
			Value:  transactionId,
			Inline: true,
		},
		{
			Name:   "Amount",
			Value:  fmt.Sprintf("%v", amount/model.OMISE_CURRENCY_RATE_TH),
			Inline: true,
		},
	}
	footer := sender.Footer{
		Text: "Ticket Transaction",
	}
	send_err := s.SendEmbedMessage(fmt.Sprintf("[Transaction: %v%v]", s.MapTransactionStatusIcon(status), status), s.MapTransactionStatusColor(status), fields, footer)
	if send_err != nil {
		return send_err
	}
	return nil
}

func (s *discordService) SendEmbedMessage(title string, color int, fields []sender.Field, footer sender.Footer) error {
	embed := s.CreateEmbed(title, color, fields, footer)
	send_err := s.SendMessage(*embed)
	if send_err != nil {
		return send_err
	}
	return nil
}

func (s *discordService) CreateEmbed(title string, color int, fields []sender.Field, footer sender.Footer) *sender.Embed {
	embed := sender.Embed{
		Title:     title,
		Color:     color,
		Timestamp: time.Now(),
		Fields:    fields,
		Footer:    footer,
	}
	return &embed
}

func (s *discordService) SendMessage(embeds sender.Embed) error {
	hook := sender.Hook{
		Username:   s.SenderName,
		Avatar_url: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		Embeds:     []sender.Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	err = sender.ExecuteWebhook(s.HookUrl, payload)
	return err
}
