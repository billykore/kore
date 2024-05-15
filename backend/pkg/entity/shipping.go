package entity

type ShippingRequest struct {
	Name string `query:"name"`
}

type ShippingResponse struct {
	Message string `json:"message"`
}