package storage

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/domain/product"
	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/domain/user"
	repo2 "github.com/billykore/kore/backend/infra/storage/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo2.NewOrderRepo, wire.Bind(new(order.Repository), new(*repo2.OrderRepo)),
	repo2.NewOtpRepo, wire.Bind(new(otp.Repository), new(*repo2.OtpRepo)),
	repo2.NewProductRepo, wire.Bind(new(product.Repository), new(*repo2.ProductRepo)),
	repo2.NewShippingRepo, wire.Bind(new(shipping.Repository), new(*repo2.ShippingRepo)),
	repo2.NewUserRepo, wire.Bind(new(user.Repository), new(*repo2.UserRepo)),
)
