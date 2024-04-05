package repo

import (
	"context"

	"github.com/billykore/kore/pkg/model"
)

type TodoRepository interface {
	Get(ctx context.Context, isDone string) ([]*model.Todo, error)
	GetById(ctx context.Context, id string) (*model.Todo, error)
	Save(ctx context.Context, todo *model.Todo) error
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}
