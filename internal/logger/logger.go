package logger

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"
)

func New(level string, w io.Writer) (zerolog.Logger, error) {
	zerolog.LevelFieldName = FieldSeverity
	zerolog.TimestampFieldName = FieldTimestamp
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logger := zerolog.New(w).With().Timestamp().Logger()

	pl, err := zerolog.ParseLevel(level)
	if err != nil {
		return logger, fmt.Errorf("[logger] failed to parse logger level: %w", err)
	}

	zerolog.SetGlobalLevel(pl)

	return logger, nil
}

func Must(level string, w io.Writer) zerolog.Logger {
	logger, err := New(level, w)
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}

	return logger
}

func WithPlace(ctx context.Context, place string) zerolog.Logger {
	return zerolog.Ctx(ctx).
		With().
		Str(FieldPlace, place).
		Logger()
}
