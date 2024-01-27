package repository

import (
	"context"

	"github.com/billykore/todolist/libs/model"
)

type User interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}
