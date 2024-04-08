package entity

type GetTodosRequest struct {
	IsDone string `query:"isDone"`
}

type GetTodosResponse []*GetTodoResponse

type GetTodoResponse struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type AddTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
