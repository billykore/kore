package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

// ProductCategoryRepository .
type ProductCategoryRepository interface {
	List(ctx context.Context) ([]*model.ProductCategory, error)
}
