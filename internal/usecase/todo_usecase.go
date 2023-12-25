package usecase

import (
	"context"

	"github.com/billykore/todolist/internal/entity"
	"github.com/billykore/todolist/internal/errors"
	"github.com/billykore/todolist/internal/pkg/log"
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

func (uc *TodoUsecase) GetTodos(ctx context.Context, param *entity.GetTodosParam) ([]*entity.Todo, error) {
	todos, err := uc.repo.GetTodos(ctx, param.IsDone)
	if err != nil {
		uc.log.Usecase("GetTodos").Error(err)
		return nil, &errors.Error{
			Type:    errors.TypeNotFound,
			Message: "Todos not found",
		}
	}
	var todosData []*entity.Todo
	for _, t := range todos {
		todosData = append(todosData, entity.NewTodo(t))
	}

	return todosData, nil
}

func (uc *TodoUsecase) GetTodo(ctx context.Context, param *entity.TodoSelectorParam) (*entity.Todo, error) {
	todo, err := uc.repo.GetTodoById(ctx, param.Id)
	if err != nil {
		uc.log.Usecase("GetTodo").Error(err)
		return nil, &errors.Error{
			Type:    errors.TypeNotFound,
			Message: "Todo not found",
		}
	}
	return entity.NewTodo(todo), nil
}

func (uc *TodoUsecase) SaveTodo(ctx context.Context, param *entity.AddTodoParam) error {
	id, err := uuid.NewUUID()
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to save todo",
		}
	}
	err = uc.repo.SaveTodo(ctx, param.ToModel(id.String()))
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to save todo",
		}
	}
	return nil
}

func (uc *TodoUsecase) SetDoneTodo(ctx context.Context, param *entity.TodoSelectorParam) error {
	err := uc.repo.SetDoneTodo(ctx, param.Id)
	if err != nil {
		uc.log.Usecase("SetDoneTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to update todo.",
		}
	}
	return nil
}

func (uc *TodoUsecase) DeleteTodo(ctx context.Context, param *entity.TodoSelectorParam) error {
	err := uc.repo.DeleteTodo(ctx, param.Id)
	if err != nil {
		uc.log.Usecase("DeleteTodo").Error(err)
		return &errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to delete todo.",
		}
	}
	return nil
}
