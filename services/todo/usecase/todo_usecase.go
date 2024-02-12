package usecase

import (
	"context"

	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/pkg/messages"
	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/libs/repo"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoUsecase struct {
	log      *log.Logger
	todoRepo repo.TodoRepository
}

func NewTodoUsecase(log *log.Logger, todoRepo repo.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		log:      log,
		todoRepo: todoRepo,
	}
}

func (uc *TodoUsecase) GetTodos(ctx context.Context, req *v1.GetTodosRequest) ([]*v1.Todo, error) {
	todos, err := uc.todoRepo.Get(ctx, req.GetIsDone())
	if err != nil {
		uc.log.Usecase("GetTodos").Error(err)
		return nil, status.Error(codes.NotFound, messages.TodosNotFound)
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
	todo, err := uc.todoRepo.GetById(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("GetTodo").Error(err)
		return nil, status.Error(codes.NotFound, messages.TodosNotFound)
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
		return status.Error(codes.Internal, messages.FailedSaveTodo)
	}
	err = uc.todoRepo.Save(ctx, &model.Todo{
		Id:          id.String(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	})
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return status.Error(codes.Internal, messages.FailedSaveTodo)
	}
	return nil
}

func (uc *TodoUsecase) SetDoneTodo(ctx context.Context, req *v1.TodoRequest) error {
	err := uc.todoRepo.Update(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("SetDoneTodo").Error(err)
		return status.Error(codes.Internal, messages.FailedSetDoneTodo)
	}
	return nil
}

func (uc *TodoUsecase) DeleteTodo(ctx context.Context, req *v1.TodoRequest) error {
	err := uc.todoRepo.Delete(ctx, req.GetId())
	if err != nil {
		uc.log.Usecase("DeleteTodo").Error(err)
		return status.Error(codes.Internal, messages.FailedSetDoneTodo)
	}
	return nil
}
