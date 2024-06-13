package handler

import (
	"net/http"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/todo/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	uc *usecase.TodoUsecase
}

func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

func (h *TodoHandler) GetTodos(ctx echo.Context) error {
	req := new(entity.GetTodosRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todos, err := h.uc.GetTodos(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("todos", todos))
}

func (h *TodoHandler) GetTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	todo, err := h.uc.GetTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("todo", todo))
}

func (h *TodoHandler) AddTodo(ctx echo.Context) error {
	req := new(entity.AddTodoRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	save, err := h.uc.SaveTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("save", save))
}

func (h *TodoHandler) SetDoneTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	setDone, err := h.uc.SetDoneTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("done", setDone))
}

func (h *TodoHandler) DeleteTodo(ctx echo.Context) error {
	req := new(entity.ParamId)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	del, err := h.uc.DeleteTodo(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("delete", del))
}
