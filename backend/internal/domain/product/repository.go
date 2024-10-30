package product

import "context"

type Repository interface {
	List(ctx context.Context, categoryId, limit, startId int) ([]*Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
	CartList(ctx context.Context, username string, limit, startId int) ([]*Cart, error)
	SaveCart(ctx context.Context, cart Cart) error
	UpdateCart(ctx context.Context, id int, cart Cart) error
	DeleteCart(ctx context.Context, id int, cart Cart) error
	CategoryList(ctx context.Context) ([]*ProductCategory, error)
	DiscountList(ctx context.Context, limit, startId int) ([]*Discount, error)
}
