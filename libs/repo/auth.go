package repo

import (
	"context"

	"github.com/billykore/kore/libs/model"
)

type AuthRepository interface {
	Login(ctx context.Context, auth *model.Auth) error
	Logout(ctx context.Context, auth *model.Auth) error
}
