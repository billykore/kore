package handler

import (
	"github.com/billykore/todolist/internal/entity"
	"github.com/billykore/todolist/internal/errors"
	"github.com/billykore/todolist/internal/pkg/api"
	"github.com/billykore/todolist/internal/pkg/log"
	"github.com/billykore/todolist/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	log *log.Logger
	uc  *usecase.TodoUsecase
}

func NewTodoHandler(log *log.Logger, uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		log: log,
		uc:  uc,
	}
}

func (h *TodoHandler) GetTodos(ctx *gin.Context) {
	param, err := h.getTodosParam(ctx)
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

func (h *TodoHandler) getTodosParam(ctx *gin.Context) (*entity.GetTodosParam, error) {
	param := new(entity.GetTodosParam)
	err := ctx.ShouldBindQuery(param)
	if err != nil {
		h.log.Error(err)
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}

func (h *TodoHandler) GetTodo(ctx *gin.Context) {
	param, err := h.todoParam(ctx)
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
	param, err := h.addTodoParam(ctx)
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

func (h *TodoHandler) addTodoParam(ctx *gin.Context) (*entity.AddTodoParam, error) {
	param := new(entity.AddTodoParam)
	err := ctx.ShouldBindJSON(param)
	if err != nil {
		h.log.Error(err)
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}

func (h *TodoHandler) SetDoneTodo(ctx *gin.Context) {
	param, err := h.todoParam(ctx)
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
	param, err := h.todoParam(ctx)
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

func (h *TodoHandler) todoParam(ctx *gin.Context) (*entity.TodoParam, error) {
	param := new(entity.TodoParam)
	err := ctx.ShouldBindUri(param)
	if err != nil {
		h.log.Error(err)
		return nil, errors.ErrInvalidRequest
	}
	return param, nil
}
