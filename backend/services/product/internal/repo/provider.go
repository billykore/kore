package repo

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewProductRepository,
	NewProductCategoryRepository,
	NewDiscountRepository,
	NewCartRepository,
)
