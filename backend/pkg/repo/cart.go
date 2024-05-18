package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type CartRepository interface {
	List(ctx context.Context, userId int, limit int, startId int) ([]*model.Cart, error)
	Save(ctx context.Context, item model.Cart) error
	Update(ctx context.Context, id int, quantity int) error
	Delete(ctx context.Context, id int) error
}
