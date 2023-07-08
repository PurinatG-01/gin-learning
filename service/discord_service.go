package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	sender "github.com/bensch777/discord-webhook-golang"
)

type DiscordService interface {
	SendTransactionMessage(transactionId string, amount int, purchaserId int, eventId int, status string) error
	SendEmbedMessage(title string, color int, fields []sender.Field, footer sender.Footer) error
	CreateEmbed(title string, color int, fields []sender.Field, footer sender.Footer) *sender.Embed
	SendMessage(embeds sender.Embed) error
}

func NewDiscordService() DiscordService {
	url := fmt.Sprintf(os.Getenv("DISCORD_WEBHOOK_URL"))
	if url == "" {
		panic("DISCORD_WEBHOOK_URL is not set")
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

func (s *discordService) SendTransactionMessage(transactionId string, amount int, purchaserId int, eventId int, status string) error {
	fields := []sender.Field{
		{
			Name:   "Transaction ID",
			Value:  transactionId,
			Inline: true,
		},
		{
			Name:   "Amount",
			Value:  fmt.Sprintf("%v", amount),
			Inline: true,
		},
		{
			Name:   "Purchaser ID",
			Value:  fmt.Sprintf("%d", purchaserId),
			Inline: true,
		},
		{
			Name:   "Event ID",
			Value:  fmt.Sprintf("%d", eventId),
			Inline: true,
		},
	}
	footer := sender.Footer{
		Text: "Ticket Transaction",
	}
	send_err := s.SendEmbedMessage("Transaction", 15277667, fields, footer)
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
		Content:    "Message",
		Embeds:     []sender.Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	err = sender.ExecuteWebhook(s.HookUrl, payload)
	return err
}
