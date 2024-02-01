package service

import (
	"context"

	v1 "github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/services/todo/usecase"
	"google.golang.org/grpc/codes"
)

type TodoService struct {
	v1.UnimplementedTodoServiceServer

	uc *usecase.TodoUsecase
}

func NewTodoService(uc *usecase.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) GetTodos(ctx context.Context, in *v1.GetTodosRequest) (*v1.GetTodosReply, error) {
	todos, err := s.uc.GetTodos(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GetTodosReply{Todos: todos}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, in *v1.TodoRequest) (*v1.GetTodoReply, error) {
	todo, err := s.uc.GetTodo(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GetTodoReply{Todo: todo}, nil
}

func (s *TodoService) AddTodo(ctx context.Context, in *v1.AddTodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.SaveTodo(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.DefaultReply{Message: codes.OK.String()}, nil
}

func (s *TodoService) SetDoneTodo(ctx context.Context, in *v1.TodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.SetDoneTodo(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.DefaultReply{Message: codes.OK.String()}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, in *v1.TodoRequest) (*v1.DefaultReply, error) {
	err := s.uc.DeleteTodo(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.DefaultReply{Message: codes.OK.String()}, nil
}
