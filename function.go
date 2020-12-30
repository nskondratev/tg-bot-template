package gcp

import (
	"context"
	"net/http"
	"os"

	"github.com/rs/zerolog"

	"github.com/nskondratev/tg-bot-template/internal/app/bot"
	"github.com/nskondratev/tg-bot-template/internal/boot"
	"github.com/nskondratev/tg-bot-template/internal/env"
	"github.com/nskondratev/tg-bot-template/internal/logger"
)

var (
	b   *bot.Bot
	log zerolog.Logger
)

func init() {
	var err error

	log = logger.Must(env.String("LOG_LEVEL", "debug"), os.Stdout)

	b, err = boot.InitBot(context.Background(), log)
	if err != nil {
		panic("failed to init bot: " + err.Error())
	}
}

func BotUpdate(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("New update")

	up, err := b.UpdateFromRequest(r)
	if err != nil {
		log.Err(err).Msg("failed to get update from request")
		http.Error(w, "bad input", http.StatusBadRequest)

		return
	}

	b.HandleUpdate(r.Context(), up)
	w.WriteHeader(http.StatusOK)
}
