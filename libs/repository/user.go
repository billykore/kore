package repository

import (
	"context"

	"github.com/billykore/kore/libs/model"
)

type User interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}
