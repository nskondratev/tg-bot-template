package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	"github.com/nskondratev/tg-bot-template/internal/metrics"
)

func InjectAndFlushMetrics(stats *metrics.Client) func(next bot.Handler) bot.Handler {
	return func(next bot.Handler) bot.Handler {
		return bot.HandlerFunc(func(ctx context.Context, bot bot.Sender, update *tgbotapi.Update) {
			ctx = metrics.WithContext(ctx, stats)

			next.Handle(ctx, bot, update)

			// Flush metrics
			stats.Flush()
		})
	}
}
