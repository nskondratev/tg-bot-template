package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
)

func Recover(next bot.Handler) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, b bot.Sender, update *tgbotapi.Update) {
		defer func() {
			if err := recover(); err != nil {
				zerolog.Ctx(ctx).
					Error().
					Interface("error", err).
					Msg("recovered from panic")
			}
		}()

		next.Handle(ctx, b, update)
	})
}
