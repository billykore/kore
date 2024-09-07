package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type CartRepository interface {
	List(ctx context.Context, username string, limit int, startId int) ([]*model.Cart, error)
	Save(ctx context.Context, cart model.Cart) error
	Update(ctx context.Context, id int, cart model.Cart) error
	Delete(ctx context.Context, id int, cart model.Cart) error
}
