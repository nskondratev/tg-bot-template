package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var NoopHandler = HandlerFunc(func(_ context.Context, _ Sender, _ *tgbotapi.Update) {})

type Sender interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type Handler interface {
	Handle(ctx context.Context, bot Sender, update *tgbotapi.Update)
}

type HandlerFunc func(ctx context.Context, bot Sender, update *tgbotapi.Update)

func (f HandlerFunc) Handle(ctx context.Context, bot Sender, update *tgbotapi.Update) {
	f(ctx, bot, update)
}

type Middleware func(next Handler) Handler

type Chain struct {
	middlewares []Middleware
}

func NewChain(middlewares ...Middleware) Chain {
	return Chain{append(([]Middleware)(nil), middlewares...)}
}

func (c Chain) Then(h Handler) Handler {
	for i := range c.middlewares {
		h = c.middlewares[len(c.middlewares)-1-i](h)
	}

	return h
}

func (c Chain) Append(middlewares ...Middleware) Chain {
	newCons := make([]Middleware, 0, len(c.middlewares)+len(middlewares))
	newCons = append(newCons, c.middlewares...)
	newCons = append(newCons, middlewares...)

	return Chain{newCons}
}

func (c Chain) Extend(nc Chain) Chain {
	return c.Append(nc.middlewares...)
}
