package entity

import (
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/types"
)

type CartRequest struct {
	Limit   int `query:"limit"`
	StartId int `query:"startId"`
}

type CartResponse struct {
	Id       uint `json:"id"`
	Quantity int  `json:"quantity"`
	Product  struct {
		Id    uint        `json:"id"`
		Name  string      `json:"name"`
		Price types.Money `json:"price"`
	} `json:"product"`
}

// MakeCartResponse makes CartResponse from model.Cart model.
func MakeCartResponse(m *model.Cart) *CartResponse {
	return &CartResponse{
		Id:       m.ID,
		Quantity: m.Quantity,
		Product: struct {
			Id    uint        `json:"id"`
			Name  string      `json:"name"`
			Price types.Money `json:"price"`
		}{
			Id:    m.Product.ID,
			Name:  m.Product.Name,
			Price: m.Product.Price,
		},
	}
}

type AddCartItemRequest struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Id       int `param:"cartId" swaggerignore:"true"`
	Quantity int `json:"quantity"`
}

type DeleteCartItemRequest struct {
	Id int `param:"cartId"`
}
