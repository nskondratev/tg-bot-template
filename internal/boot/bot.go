package boot

import (
	"context"
	"os"

	"github.com/rs/zerolog"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	"github.com/nskondratev/tg-bot-template/internal/app/bot/handlers/command"
	"github.com/nskondratev/tg-bot-template/internal/app/bot/handlers/message"
	"github.com/nskondratev/tg-bot-template/internal/app/bot/middleware"
)

func InitBot(_ context.Context, log zerolog.Logger) (*bot.Bot, error) {
	// Here you can create clients to external services, databases, etc.
	// Set up bot handlers
	updateHandler := bot.
		NewChain(
			command.NewHandler().Middleware,
			message.NewHandler().Middleware,
		).
		Then(bot.NoopHandler)

	handler := bot.
		NewChain(
			middleware.NonNilUpdate,
			middleware.InjectLogger(log),
			middleware.Recover,
			middleware.SetUser,
			middleware.LogUserInfo,
		).
		Then(updateHandler)

	b := bot.Must(os.Getenv("TELEGRAM_API_TOKEN"), handler)

	return b, nil
}
