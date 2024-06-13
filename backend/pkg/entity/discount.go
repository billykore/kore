package entity

import "github.com/billykore/kore/backend/pkg/model"

type DiscountRequest struct {
	Limit   int `query:"limit"`
	StartId int `query:"startId"`
}

type DiscountResponse struct {
	Id                 uint    `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	DiscountPercentage float64 `json:"discountPercentage"`
}

// MakeDiscountResponse makes DiscountResponse from model.Discount model.
func MakeDiscountResponse(m *model.Discount) *DiscountResponse {
	return &DiscountResponse{
		Id:                 m.ID,
		Name:               m.Name,
		Description:        m.Description,
		DiscountPercentage: m.DiscountPercentage,
	}
}
