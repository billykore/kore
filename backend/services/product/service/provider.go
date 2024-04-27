package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewProductService,
	NewProductCategoryService,
	NewDiscountService,
	NewCartService,
)
