package telegrambot

import (
	"context"

	"github.com/go-telegram/bot"
)

type BotTg struct {
	Bot    *bot.Bot
	Ctx    context.Context
	ChatID int
}

func BotInit(botToken string, chatID int) BotTg {
	ctx := context.Background()
	Bot, _ := bot.New(botToken)
	return BotTg{Bot, ctx, chatID}
}

func (botTg *BotTg) SendTelegramMessage(text string) error {
	_, err := botTg.Bot.SendMessage(botTg.Ctx, &bot.SendMessageParams{
		ChatID: botTg.ChatID,
		Text:   text})
	return err
}
