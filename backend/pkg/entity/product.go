package entity

import "github.com/billykore/kore/backend/pkg/model"

type ProductRequest struct {
	Name string `query:"name"`
}

type ProductResponse struct {
	Name             string                   `json:"name"`
	Description      string                   `json:"description"`
	SKU              string                   `json:"sku"`
	Price            uint64                   `json:"price"`
	CategoryId       int                      `json:"categoryId"`
	InventoryId      int                      `json:"inventoryId"`
	DiscountId       int                      `json:"discountId"`
	ProductCategory  ProductCategoryResponse  `json:"productCategory"`
	ProductInventory ProductInventoryResponse `json:"productInventory"`
	Discount         DiscountResponse         `json:"discount"`
}

func MakeProductResponse(m model.Product) ProductResponse {
	return ProductResponse{
		Name:             m.Name,
		Description:      m.Description,
		SKU:              m.SKU,
		Price:            m.Price,
		CategoryId:       m.CategoryId,
		InventoryId:      m.InventoryId,
		DiscountId:       m.DiscountId,
		ProductCategory:  MakeProductCategoryResponse(m.ProductCategory),
		ProductInventory: MakeProductInventoryResponse(m.ProductInventory),
		Discount:         MakeDiscountResponse(m.Discount),
	}
}

type ProductCategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func MakeProductCategoryResponse(m model.ProductCategory) ProductCategoryResponse {
	return ProductCategoryResponse{
		Name:        m.Name,
		Description: m.Description,
	}
}

type ProductInventoryResponse struct {
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

func MakeProductInventoryResponse(m model.ProductInventory) ProductInventoryResponse {
	return ProductInventoryResponse{
		Quantity:    m.Quantity,
		Description: m.Description,
	}
}

type DiscountResponse struct {
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	DiscountPercentage float64 `json:"discountPercentage"`
}

func MakeDiscountResponse(m model.Discount) DiscountResponse {
	return DiscountResponse{
		Name:               m.Name,
		Description:        m.Description,
		DiscountPercentage: m.DiscountPercentage,
	}
}
