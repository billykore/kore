package entity

import (
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/types"
)

type ProductRequest struct {
	ProductId  int `param:"productId"`
	CategoryId int `query:"categoryId"`
	Limit      int `query:"limit"`
	StartId    int `query:"startId"`
}

type ProductResponse struct {
	Id               uint                      `json:"id"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	SKU              string                    `json:"sku"`
	Price            types.Money               `json:"price"`
	CategoryId       int                       `json:"categoryId"`
	InventoryId      int                       `json:"inventoryId"`
	DiscountId       int                       `json:"discountId"`
	ProductCategory  *ProductCategoryResponse  `json:"productCategory"`
	ProductInventory *ProductInventoryResponse `json:"productInventory"`
	Discount         *DiscountResponse         `json:"discount"`
}

func MakeProductResponse(m *model.Product) *ProductResponse {
	return &ProductResponse{
		Id:               m.ID,
		Name:             m.Name,
		Description:      m.Description,
		SKU:              m.SKU,
		Price:            m.Price,
		CategoryId:       m.CategoryId,
		InventoryId:      m.InventoryId,
		DiscountId:       m.DiscountId,
		ProductCategory:  MakeProductCategoryResponse(&m.ProductCategory),
		ProductInventory: MakeProductInventoryResponse(&m.ProductInventory),
		Discount:         MakeDiscountResponse(&m.Discount),
	}
}
