package entity

import "github.com/billykore/kore/backend/pkg/model"

type ProductCategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MakeProductCategoryResponse makes ProductCategoryResponse from model.ProductCategory model.
func MakeProductCategoryResponse(m *model.ProductCategory) *ProductCategoryResponse {
	return &ProductCategoryResponse{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
	}
}
