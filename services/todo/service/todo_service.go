package service

import (
	"github.com/billykore/kore/pkg/entity"
	"github.com/billykore/kore/services/todo/usecase"
	"github.com/labstack/echo/v4"
)

type TodoService struct {
	uc *usecase.TodoUsecase
}

func NewTodoService(uc *usecase.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) GetTodos(ctx echo.Context) error {
	req := new(entity.GetTodosRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todos, err := s.uc.GetTodos(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(todos))
}

func (s *TodoService) GetTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todo, err := s.uc.GetTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(todo))
}

func (s *TodoService) AddTodo(ctx echo.Context) error {
	req := new(entity.AddTodoRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	save, err := s.uc.SaveTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(save))
}

func (s *TodoService) SetDoneTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	setDone, err := s.uc.SetDoneTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(setDone))
}

func (s *TodoService) DeleteTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	del, err := s.uc.DeleteTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(del))
}
