package entity

import (
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/types"
)

type CartRequest struct {
	UserId  int `query:"userId"`
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
	UserId    int
	ProductId int
	Quantity  int
}

type UpdateCartItemRequest struct {
	Id       int `param:"cartId"`
	Quantity int
}

type DeleteCartItemRequest struct {
	Id int `param:"cartId"`
}
