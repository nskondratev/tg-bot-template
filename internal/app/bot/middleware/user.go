package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"

	"github.com/nskondratev/tg-bot-template/internal/app"
	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	l "github.com/nskondratev/tg-bot-template/internal/logger"
)

type sender struct {
	chatID int64
	user   *tgbotapi.User
}

func SetUser(next bot.Handler) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, bot bot.Sender, update *tgbotapi.Update) {
		sender := getSenderFromUpdate(update)
		if sender != nil {
			u := app.User{
				TelegramUserID: int64(sender.user.ID),
				UserName:       sender.user.UserName,
				FirstName:      sender.user.FirstName,
				LastName:       sender.user.LastName,
				Lang:           sender.user.LanguageCode,
			}

			hLog := zerolog.Ctx(ctx).With().
				Int64(l.FieldUserID, u.TelegramUserID).
				Str(l.FieldUsername, u.UserName).
				Logger()

			ctx = hLog.WithContext(ctx)
			ctx = app.NewContextWithUser(ctx, u)
		}

		next.Handle(ctx, bot, update)
	})
}

func getSenderFromUpdate(update *tgbotapi.Update) *sender {
	switch {
	case update.Message != nil && update.Message.From != nil:
		return &sender{
			chatID: update.Message.Chat.ID,
			user:   update.Message.From,
		}
	case update.CallbackQuery != nil && update.CallbackQuery.From != nil:
		return &sender{
			chatID: update.CallbackQuery.Message.Chat.ID,
			user:   update.CallbackQuery.From,
		}
	default:
		return nil
	}
}
