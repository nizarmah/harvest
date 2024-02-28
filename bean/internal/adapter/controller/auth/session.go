package auth

import (
	"context"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

type authContextKey string

const sessionContextKey = authContextKey("session")

func NewContextWithSession(ctx context.Context, session *model.Session) context.Context {
	return context.WithValue(ctx, sessionContextKey, session)
}

func SessionFromContext(ctx context.Context) *model.Session {
	session, ok := ctx.Value(sessionContextKey).(*model.Session)
	if !ok {
		return nil
	}

	return session
}
