package repo

import (
	"github.com/billykore/kore/backend/pkg/model"
	"golang.org/x/net/context"
)

type ProductRepository interface {
	List(ctx context.Context) ([]model.Product, error)
}
