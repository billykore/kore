package repo

import (
	"github.com/billykore/kore/backend/pkg/model"
	"golang.org/x/net/context"
)

// ProductRepository .
type ProductRepository interface {
	List(ctx context.Context, categoryId int, limit int, startId int) ([]*model.Product, error)
	GetById(ctx context.Context, id int) (*model.Product, error)
}
