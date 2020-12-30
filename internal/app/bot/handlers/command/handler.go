package command

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	l "github.com/nskondratev/tg-bot-template/internal/logger"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Middleware(next bot.Handler) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, bot bot.Sender, update *tgbotapi.Update) {
		if update.Message != nil && update.Message.IsCommand() {
			h.Handle(ctx, bot, update)

			return
		}

		next.Handle(ctx, bot, update)
	})
}

func (h *Handler) Handle(ctx context.Context, bot bot.Sender, update *tgbotapi.Update) {
	if update.Message == nil || !update.Message.IsCommand() {
		return
	}

	cmd := update.Message.Command()

	log := l.WithPlace(ctx, "bot_handlers_command")

	log.Info().
		Str(l.FieldCommand, cmd).
		Msg("command received")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Command received: %s", cmd))
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to send message")
	}
}
