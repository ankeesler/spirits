package rest

import "context"

var restPathContextKey struct{}

func WithPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, restPathContextKey, path)
}

func Path(ctx context.Context) string {
	return ctx.Value(restPathContextKey).(string)
}
