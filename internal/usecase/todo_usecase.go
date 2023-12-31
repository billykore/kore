package usecase

import (
	"context"

	"github.com/billykore/todolist/internal/errors"
	"github.com/billykore/todolist/internal/model"
	"github.com/billykore/todolist/internal/pkg/log"
	v1 "github.com/billykore/todolist/internal/proto/v1"
	"github.com/billykore/todolist/internal/repository"
	"github.com/google/uuid"
)

type TodoUsecase struct {
	log  *log.Logger
	repo *repository.TodoRepository
}

func NewTodoUsecase(log *log.Logger, repo *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *TodoUsecase) GetTodos(ctx context.Context, req *v1.GetTodosRequest) ([]*v1.Todo, error) {
	todos, err := uc.repo.GetTodos(ctx, req.GetIsDone())
	if err != nil {
		uc.log.Usecase("GetTodos").Error(err)
		return nil, &errors.Error{
			Type:    errors.TypeNotFound,
			Message: "Todos not found",
		}
	}
	var todosData []*v1.Todo
	for _, t := range todos {
		todosData = append(todosData, &v1.Todo{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
		})
	}
	return todosData, nil
}

func (uc *TodoUsecase) GetTodo(ctx context.Context, req *v1.TodoRequest) (*v1.Todo, error) {
	todo, err := uc.repo.GetTodoById(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("GetTodo").Error(err)
		return nil, &errors.Error{
			Type:    errors.TypeNotFound,
			Message: "Todo not found",
		}
	}
	return &v1.Todo{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}, nil
}

func (uc *TodoUsecase) SaveTodo(ctx context.Context, req *v1.AddTodoRequest) error {
	id, err := uuid.NewUUID()
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to save todo",
		}
	}
	err = uc.repo.SaveTodo(ctx, &model.Todo{
		Id:          id.String(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	})
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to save todo",
		}
	}
	return nil
}

func (uc *TodoUsecase) SetDoneTodo(ctx context.Context, req *v1.TodoRequest) error {
	err := uc.repo.SetDoneTodo(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("SetDoneTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to update todo",
		}
	}
	return nil
}

func (uc *TodoUsecase) DeleteTodo(ctx context.Context, req *v1.TodoRequest) error {
	err := uc.repo.DeleteTodo(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("DeleteTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to delete todo",
		}
	}
	return nil
}
