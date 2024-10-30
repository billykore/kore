package product

import (
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/pkg/types"
)

type GetRequest struct {
	ProductId  int `param:"productId"`
	CategoryId int `query:"categoryId"`
	Limit      int `query:"limit"`
	StartId    int `query:"startId"`
}

type GetResponse struct {
	Id               uint               `json:"id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	SKU              string             `json:"sku"`
	Price            types.Money        `json:"price"`
	CategoryId       int                `json:"categoryId"`
	InventoryId      int                `json:"inventoryId"`
	DiscountId       int                `json:"discountId"`
	ProductCategory  *CategoryResponse  `json:"productCategory"`
	ProductInventory *InventoryResponse `json:"productInventory"`
	Discount         *DiscountResponse  `json:"discount"`
}

// MakeResponse makes ProductResponse from model.Order model.
func MakeResponse(m *product.Product) *GetResponse {
	return &GetResponse{
		Id:               m.ID,
		Name:             m.Name,
		Description:      m.Description,
		SKU:              m.SKU,
		Price:            m.Price,
		CategoryId:       m.CategoryId,
		InventoryId:      m.InventoryId,
		DiscountId:       m.DiscountId,
		ProductCategory:  MakeCategoryResponse(&m.Category),
		ProductInventory: MakeInventoryResponse(&m.Inventory),
		Discount:         MakeDiscountResponse(&m.Discount),
	}
}

type InventoryResponse struct {
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

// MakeInventoryResponse makes InventoryResponse from model.ProductInventory model.
func MakeInventoryResponse(m *product.Inventory) *InventoryResponse {
	return &InventoryResponse{
		Quantity:    m.Quantity,
		Description: m.Description,
	}
}

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
func MakeCartResponse(m *product.Cart) *CartResponse {
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
func MakeDiscountResponse(m *product.Discount) *DiscountResponse {
	return &DiscountResponse{
		Id:                 m.ID,
		Name:               m.Name,
		Description:        m.Description,
		DiscountPercentage: m.DiscountPercentage,
	}
}

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MakeCategoryResponse makes CategoryResponse from model.ProductCategory model.
func MakeCategoryResponse(m *product.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
	}
}
