package handler

import (
	"github.com/billykore/todolist/internal/entity"
	"github.com/billykore/todolist/internal/errors"
	"github.com/billykore/todolist/internal/pkg/api"
	"github.com/billykore/todolist/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	uc *usecase.TodoUsecase
}

func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		uc: uc,
	}
}

func (h *TodoHandler) GetTodos(ctx *gin.Context) {
	param, err := getTodosParam(ctx)
	if err != nil {
		ctx.JSON(api.ResponseBadRequest(err))
		return
	}
	todos, err := h.uc.GetTodos(ctx, param)
	if err != nil {
		ctx.JSON(api.ResponseError(err))
		return
	}
	ctx.JSON(api.ResponseSuccess(todos))
}

func getTodosParam(ctx *gin.Context) (*entity.GetTodosParam, error) {
	param := new(entity.GetTodosParam)
	err := ctx.ShouldBindQuery(param)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}

func (h *TodoHandler) GetTodo(ctx *gin.Context) {
	param, err := todoSelectorParam(ctx)
	if err != nil {
		ctx.JSON(api.ResponseBadRequest(err))
		return
	}
	todo, err := h.uc.GetTodo(ctx, param)
	if err != nil {
		ctx.JSON(api.ResponseError(err))
		return
	}
	ctx.JSON(api.ResponseSuccess(todo))
}

func (h *TodoHandler) AddTodo(ctx *gin.Context) {
	param, err := addTodoParam(ctx)
	if err != nil {
		ctx.JSON(api.ResponseBadRequest(err))
		return
	}
	err = h.uc.SaveTodo(ctx, param)
	if err != nil {
		ctx.JSON(api.ResponseError(err))
		return
	}
	ctx.JSON(api.ResponseSuccess(nil))
}

func addTodoParam(ctx *gin.Context) (*entity.AddTodoParam, error) {
	param := new(entity.AddTodoParam)
	err := ctx.ShouldBindJSON(param)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}

func (h *TodoHandler) SetDoneTodo(ctx *gin.Context) {
	param, err := todoSelectorParam(ctx)
	if err != nil {
		ctx.JSON(api.ResponseBadRequest(err))
		return
	}
	err = h.uc.SetDoneTodo(ctx, param)
	if err != nil {
		ctx.JSON(api.ResponseError(err))
		return
	}
	ctx.JSON(api.ResponseSuccess(nil))
}

func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	param, err := todoSelectorParam(ctx)
	if err != nil {
		ctx.JSON(api.ResponseBadRequest(err))
		return
	}
	err = h.uc.DeleteTodo(ctx, param)
	if err != nil {
		ctx.JSON(api.ResponseError(err))
		return
	}
	ctx.JSON(api.ResponseSuccess(nil))
}

func todoSelectorParam(ctx *gin.Context) (*entity.TodoSelectorParam, error) {
	param := new(entity.TodoSelectorParam)
	err := ctx.ShouldBindUri(param)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}
