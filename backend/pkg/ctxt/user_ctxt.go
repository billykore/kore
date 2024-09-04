package ctxt

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
)

const UserKey = "user"

func ContextWithUser(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func UserFromContext(ctx context.Context) (entity.User, bool) {
	user, ok := ctx.Value(UserKey).(entity.User)
	return user, ok
}
