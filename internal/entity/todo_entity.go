package entity

import "github.com/billykore/todolist/internal/model"

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

func NewTodo(model *model.Todo) *Todo {
	return &Todo{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		IsDone:      model.IsDone,
	}
}

type GetTodosParam struct {
	IsDone string `form:"isDone"`
}

type TodoParam struct {
	Id string `uri:"id"`
}

type AddTodoParam struct {
	Title       string
	Description string
}

func (param *AddTodoParam) ToModel(id string) *model.Todo {
	return &model.Todo{
		Id:          id,
		Title:       param.Title,
		Description: param.Description,
		IsDone:      false,
	}
}
