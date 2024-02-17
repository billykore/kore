package entity

type BookRequest struct {
	Name string `query:"name"`
}

type BookResponse struct {
	Message string `json:"message"`
}