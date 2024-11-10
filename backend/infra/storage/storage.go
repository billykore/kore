package storage

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/domain/product"
	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/domain/user"
	"github.com/billykore/kore/backend/infra/storage/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewOrderRepo, wire.Bind(new(order.Repository), new(*repo.OrderRepo)),
	repo.NewOtpRepo, wire.Bind(new(otp.Repository), new(*repo.OtpRepo)),
	repo.NewProductRepo, wire.Bind(new(product.Repository), new(*repo.ProductRepo)),
	repo.NewShippingRepo, wire.Bind(new(shipping.Repository), new(*repo.ShippingRepo)),
	repo.NewUserRepo, wire.Bind(new(user.Repository), new(*repo.UserRepo)),
)
