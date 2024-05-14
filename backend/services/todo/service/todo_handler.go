package service

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/todo/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	uc *usecase.TodoUsecase
}

func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (s *TodoHandler) GetTodos(ctx echo.Context) error {
	req := new(entity.GetTodosRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todos, err := s.uc.GetTodos(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("todos", todos))
}

func (s *TodoHandler) GetTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todo, err := s.uc.GetTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("todo", todo))
}

func (s *TodoHandler) AddTodo(ctx echo.Context) error {
	req := new(entity.AddTodoRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	save, err := s.uc.SaveTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("save", save))
}

func (s *TodoHandler) SetDoneTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	setDone, err := s.uc.SetDoneTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("done", setDone))
}

func (s *TodoHandler) DeleteTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	del, err := s.uc.DeleteTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("delete", del))
}
