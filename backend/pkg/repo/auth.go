package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type AuthRepository interface {
	Login(ctx context.Context, auth *model.AuthActivities) error
	Logout(ctx context.Context, auth *model.AuthActivities) error
}
