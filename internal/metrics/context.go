package metrics

import "context"

type metricsContextKey struct{}

func WithContext(ctx context.Context, metrics *Client) context.Context {
	return context.WithValue(ctx, metricsContextKey{}, metrics)
}

func FromContext(ctx context.Context) *Client {
	m, ok := ctx.Value(metricsContextKey{}).(*Client)
	if !ok || m == nil {
		return &Client{}
	}

	return m
}
