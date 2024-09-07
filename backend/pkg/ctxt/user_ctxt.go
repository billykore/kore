package ctxt

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/pkg/entity"
)

var ErrGetUserFromContext = errors.New("failed to get user from context")

const UserKey = "user"

// ContextWithUser set user data to the ctx context.
func ContextWithUser(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

// UserFromContext gets user data from ctx context.
func UserFromContext(ctx context.Context) (entity.User, bool) {
	user, ok := ctx.Value(UserKey).(entity.User)
	return user, ok
}
