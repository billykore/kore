package entity

import "github.com/billykore/kore/backend/pkg/model"

type ProductInventoryResponse struct {
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

func MakeProductInventoryResponse(m *model.ProductInventory) *ProductInventoryResponse {
	return &ProductInventoryResponse{
		Quantity:    m.Quantity,
		Description: m.Description,
	}
}
