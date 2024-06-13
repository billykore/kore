package repo

import (
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(repo.ProductRepository), new(*ProductRepo)),
	NewProductRepository,

	wire.Bind(new(repo.ProductCategoryRepository), new(*ProductCategoryRepo)),
	NewProductCategoryRepository,

	wire.Bind(new(repo.DiscountRepository), new(*DiscountRepo)),
	NewDiscountRepository,

	wire.Bind(new(repo.CartRepository), new(*CartRepo)),
	NewCartRepository,
)
