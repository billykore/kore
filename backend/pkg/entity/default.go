package entity

type ParamId struct {
	Id int64 `param:"id"`
}

type Message struct {
	Message string `json:"message"`
}
