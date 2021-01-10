package metrics

import (
	"context"
	"fmt"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/rs/zerolog"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	l "github.com/nskondratev/tg-bot-template/internal/logger"
)

var (
	// Measures
	commandsCount = stats.Int64("command_call", "Number of command calls", "Count")
	messagesCount = stats.Int64("messages", "Number of incoming messages", "Count")

	// Tags
	keyCommand, _      = tag.NewKey("command")
	keyActionResult, _ = tag.NewKey("action_result")
)

func views() []*view.View {
	return []*view.View{
		{
			Name:        "commands_count",
			Description: "The number of incoming commands",
			TagKeys:     []tag.Key{keyCommand, keyActionResult},
			Measure:     commandsCount,
			Aggregation: view.Count(),
		},
		{
			Name:        "messages_count",
			Description: "The number of incoming messages",
			TagKeys:     []tag.Key{keyActionResult},
			Measure:     messagesCount,
			Aggregation: view.Count(),
		},
	}
}

type Client struct {
	enabled  bool
	exporter *stackdriver.Exporter
	log      zerolog.Logger
}

func New(ctx context.Context) (*Client, error) {
	c := &Client{log: l.WithPlace(ctx, "metrics"), enabled: true}

	err := view.Register(views()...)
	if err != nil {
		return nil, fmt.Errorf("[metrics] failed to register views: %w", err)
	}

	c.exporter, err = stackdriver.NewExporter(stackdriver.Options{})
	if err != nil {
		return nil, fmt.Errorf("[metrics] failed to create exporter: %w", err)
	}

	if err := c.exporter.StartMetricsExporter(); err != nil {
		return nil, fmt.Errorf("[metrics] failed to start exporter: %w", err)
	}

	// Wait for context canceling
	go func() {
		<-ctx.Done()

		c.log.Info().Msg("metrics client context is closed. Stop metrics exporter")

		c.exporter.Flush()
		c.exporter.StopMetricsExporter()
	}()

	return c, nil
}

// Flush waits for all metrics to be exported
func (c *Client) Flush() {
	if !c.enabled {
		return
	}

	c.exporter.Flush()
}

func (c *Client) RecordCommand(ctx context.Context, command, actionResult string) {
	if !c.enabled {
		return
	}

	ctx, err := tag.New(ctx, tag.Insert(keyCommand, command), tag.Insert(keyActionResult, actionResult))
	if err != nil {
		c.log.Warn().
			Err(err).
			Msg("failed to create context with tags")

		return
	}

	stats.Record(ctx, commandsCount.M(1))
}

func (c *Client) RecordMessage(ctx context.Context, actionResult string) {
	if !c.enabled {
		return
	}

	ctx, err := tag.New(ctx, tag.Insert(keyActionResult, actionResult))
	if err != nil {
		c.log.Warn().
			Err(err).
			Msg("failed to create context with tags")

		return
	}

	stats.Record(ctx, messagesCount.M(1))
}
