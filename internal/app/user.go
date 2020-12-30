package app

import "context"

type User struct {
	TelegramUserID int64
	UserName       string
	FirstName      string
	LastName       string
	Lang           string
}

type userCtxKey struct{}

func UserFromContext(ctx context.Context) (user User, ok bool) {
	user, ok = ctx.Value(userCtxKey{}).(User)

	return
}

func NewContextWithUser(ctx context.Context, user User) context.Context {
	if u, ok := ctx.Value(userCtxKey{}).(User); ok {
		if u == user {
			return ctx
		}
	}

	return context.WithValue(ctx, userCtxKey{}, user)
}
