package product

import (
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

// makeResponse makes ProductResponse from model.Order model.
func makeResponse(p *Product) *GetResponse {
	return &GetResponse{
		Id:               p.ID,
		Name:             p.Name,
		Description:      p.Description,
		SKU:              p.SKU,
		Price:            p.Price,
		CategoryId:       p.CategoryId,
		InventoryId:      p.InventoryId,
		DiscountId:       p.DiscountId,
		ProductCategory:  makeCategoryResponse(&p.Category),
		ProductInventory: makeInventoryResponse(&p.Inventory),
		Discount:         makeDiscountResponse(&p.Discount),
	}
}

type InventoryResponse struct {
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

// makeInventoryResponse makes InventoryResponse from model.ProductInventory model.
func makeInventoryResponse(i *Inventory) *InventoryResponse {
	return &InventoryResponse{
		Quantity:    i.Quantity,
		Description: i.Description,
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

// makeCartResponse makes CartResponse from model.Cart model.
func makeCartResponse(c *Cart) *CartResponse {
	return &CartResponse{
		Id:       c.ID,
		Quantity: c.Quantity,
		Product: struct {
			Id    uint        `json:"id"`
			Name  string      `json:"name"`
			Price types.Money `json:"price"`
		}{
			Id:    c.Product.ID,
			Name:  c.Product.Name,
			Price: c.Product.Price,
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

// makeDiscountResponse makes DiscountResponse from model.Discount model.
func makeDiscountResponse(d *Discount) *DiscountResponse {
	return &DiscountResponse{
		Id:                 d.ID,
		Name:               d.Name,
		Description:        d.Description,
		DiscountPercentage: d.DiscountPercentage,
	}
}

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// makeCategoryResponse makes CategoryResponse from model.ProductCategory model.
func makeCategoryResponse(c *Category) *CategoryResponse {
	return &CategoryResponse{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
