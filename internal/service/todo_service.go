package service

import (
	"context"

	"github.com/billykore/todolist/internal/entity"
	"github.com/billykore/todolist/internal/errors"
	v1 "github.com/billykore/todolist/internal/grpc/v1"
	"github.com/billykore/todolist/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoService struct {
	v1.UnimplementedTodoServiceServer

	uc *usecase.TodoUsecase
}

func NewTodoService(uc *usecase.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) GetTodos(ctx context.Context, in *v1.GetTodosRequest) (*v1.GetTodosReply, error) {
	todos, err := s.uc.GetTodos(ctx, &entity.GetTodosParam{IsDone: in.IsDone})
	if err != nil {
		return &v1.GetTodosReply{}, errors.Error{
			Type:    errors.TypeNotFound,
			Message: "Todos not found",
		}
	}

	var todosMsg []*v1.Todo
	for _, t := range todos {
		todosMsg = append(todosMsg, t.GRPCMessage())
	}

	return &v1.GetTodosReply{Todos: todosMsg}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, in *v1.TodoRequest) (*v1.GetTodoReply, error) {
	todo, err := s.uc.GetTodo(ctx, &entity.TodoParam{Id: in.Id})
	if err != nil {
		return &v1.GetTodoReply{}, status.Error(codes.NotFound, "Todo not found")
	}
	return &v1.GetTodoReply{Todo: todo.GRPCMessage()}, nil
}

func (s *TodoService) AddTodo(ctx context.Context, in *v1.AddTodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.SaveTodo(ctx, &entity.AddTodoParam{
		Title:       in.Title,
		Description: in.Description,
	})
	if err != nil {
		return &v1.DefaultReply{}, errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to add todo",
		}
	}
	return &v1.DefaultReply{Message: "SUCCESS"}, nil
}

func (s *TodoService) SetDoneTodo(ctx context.Context, in *v1.TodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.SetDoneTodo(ctx, &entity.TodoParam{Id: in.Id})
	if err != nil {
		return &v1.DefaultReply{}, errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to set done todo",
		}
	}
	return &v1.DefaultReply{Message: "SUCCESS"}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, in *v1.TodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.DeleteTodo(ctx, &entity.TodoParam{Id: in.Id})
	if err != nil {
		return &v1.DefaultReply{}, errors.Error{
			Type:    errors.TypeInternalServerError,
			Message: "Failed to set done todo",
		}
	}
	return &v1.DefaultReply{Message: "SUCCESS"}, nil
}
