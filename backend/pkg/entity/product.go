package entity

type ProductRequest struct {
	Name string `query:"name"`
}

type ProductResponse struct {
	Message string `json:"message"`
}