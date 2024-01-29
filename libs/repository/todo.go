package repository

import (
	"context"

	"github.com/billykore/todolist/libs/model"
)

type Todo interface {
	GetTodos(ctx context.Context, isDone string) ([]*model.Todo, error)
	GetTodoById(ctx context.Context, id string) (*model.Todo, error)
	SaveTodo(ctx context.Context, todo *model.Todo) error
	UpdateTodo(ctx context.Context, id string) error
	DeleteTodo(ctx context.Context, id string) error
}
