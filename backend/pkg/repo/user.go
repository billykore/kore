package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}
