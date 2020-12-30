package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
)

func NonNilUpdate(next bot.Handler) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, b bot.Sender, update *tgbotapi.Update) {
		if update == nil {
			return
		}

		next.Handle(ctx, b, update)
	})
}
