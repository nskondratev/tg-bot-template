package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	l "github.com/nskondratev/tg-bot-template/internal/logger"
)

func InjectLogger(log zerolog.Logger) func(next bot.Handler) bot.Handler {
	return func(next bot.Handler) bot.Handler {
		return bot.HandlerFunc(func(ctx context.Context, b bot.Sender, update *tgbotapi.Update) {
			hLog := log
			if update != nil {
				hLog = hLog.With().Int(l.FieldUpdateID, update.UpdateID).Logger()
			}
			ctx = hLog.WithContext(ctx)

			next.Handle(ctx, b, update)
		})
	}
}

func LogUserInfo(next bot.Handler) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, b bot.Sender, update *tgbotapi.Update) {
		zerolog.Ctx(ctx).Debug().Msg("new update from user")

		next.Handle(ctx, b, update)
	})
}
