package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/messages"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/google/uuid"
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

func (uc *TodoUsecase) GetTodos(ctx context.Context, req *entity.GetTodosRequest) (entity.GetTodosResponse, error) {
	todos, err := uc.todoRepo.Get(ctx, req.IsDone)
	if err != nil {
		uc.log.Usecase("GetTodos").Error(err)
		return nil, status.Error(codes.NotFound, messages.TodosNotFound)
	}
	var todosData entity.GetTodosResponse
	for _, t := range todos {
		todosData = append(todosData, &entity.GetTodoResponse{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
		})
	}
	return todosData, nil
}

func (uc *TodoUsecase) GetTodo(ctx context.Context, req *entity.ParamId) (*entity.GetTodoResponse, error) {
	todo, err := uc.todoRepo.GetById(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("GetTodo").Error(err)
		return nil, status.Error(codes.NotFound, messages.TodosNotFound)
	}
	return &entity.GetTodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}, nil
}

func (uc *TodoUsecase) SaveTodo(ctx context.Context, req *entity.AddTodoRequest) (*entity.Message, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return nil, status.Error(codes.Internal, messages.FailedSaveTodo)
	}
	err = uc.todoRepo.Save(ctx, &model.Todo{
		Id:          id.String(),
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		uc.log.Usecase("SaveTodo").Error(err)
		return nil, status.Error(codes.Internal, messages.FailedSaveTodo)
	}
	return &entity.Message{Message: messages.SuccessSaveTodo}, nil
}

func (uc *TodoUsecase) SetDoneTodo(ctx context.Context, req *entity.ParamId) (*entity.Message, error) {
	err := uc.todoRepo.Update(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("SetDoneTodo").Error(err)
		return nil, status.Error(codes.Internal, messages.FailedSetDoneTodo)
	}
	return &entity.Message{Message: messages.SuccessSetDoneTodo}, nil
}

func (uc *TodoUsecase) DeleteTodo(ctx context.Context, req *entity.ParamId) (*entity.Message, error) {
	err := uc.todoRepo.Delete(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("DeleteTodo").Error(err)
		return nil, status.Error(codes.Internal, messages.FailedDeleteTodo)
	}
	return &entity.Message{Message: messages.SuccessDeleteTodo}, nil
}
