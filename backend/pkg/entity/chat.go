package entity

type ChatRequest struct {
	Name string `query:"name"`
}

type ChatResponse struct {
	Message string `json:"message"`
}