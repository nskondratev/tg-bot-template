package bot

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	tg      *tgbotapi.BotAPI
	handler Handler
}

func New(apiToken string, handler Handler) (*Bot, error) {
	if apiToken == "" {
		return nil, errors.New("[app_bot] api token must be provided")
	}

	if handler == nil {
		return nil, errors.New("[app_bot] handler must be provided")
	}

	tg, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, errors.New("[app_bot] failed to create telegram bot instance")
	}

	return &Bot{tg: tg, handler: handler}, nil
}

func Must(apiToken string, handler Handler) *Bot {
	b, err := New(apiToken, handler)
	if err != nil {
		panic(err)
	}

	return b
}

func (b *Bot) PollUpdates(ctx context.Context) error {
	if b.handler == nil {
		panic("handler must be set before running updater")
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := b.tg.GetUpdatesChan(updateConfig)
	if err != nil {
		return fmt.Errorf("failed to get updates channel: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update, ok := <-updates:
			if !ok {
				return nil
			}

			b.handler.Handle(ctx, b, &update)
		}
	}
}

func (b *Bot) UpdateFromRequest(r *http.Request) (*tgbotapi.Update, error) {
	return b.tg.HandleUpdate(r)
}

func (b *Bot) HandleUpdate(ctx context.Context, update *tgbotapi.Update) {
	if update == nil {
		return
	}

	b.handler.Handle(ctx, b, update)
}

func (b *Bot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return b.tg.Send(c)
}

func (b *Bot) GetFileDirectURL(fileID string) (string, error) {
	return b.tg.GetFileDirectURL(fileID)
}
