package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"

	"github.com/nskondratev/tg-bot-template/internal/boot"
	"github.com/nskondratev/tg-bot-template/internal/env"
	"github.com/nskondratev/tg-bot-template/internal/logger"
)

func main() {
	_ = godotenv.Load()

	log := logger.Must(env.String("LOG_LEVEL", "debug"), os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())

	b, err := boot.InitBot(ctx, log)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init bot")
	}

	g, ctx := errgroup.WithContext(ctx)

	// Wait for interruption
	g.Go(func() error {
		ic := make(chan os.Signal, 1)
		signal.Notify(ic, os.Interrupt, syscall.SIGTERM)
		<-ic
		log.Info().Msg("application is interrupted. Stopping appCtx...")
		cancel()

		return ctx.Err()
	})

	// Poll bot updates
	g.Go(func() error {
		log.Info().Msg("Starting polling updates...")

		return b.PollUpdates(ctx)
	})

	err = g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal().Err(err).Msg("errgroup finished with error")
	}

	log.Info().Msg("exit from app")
}
