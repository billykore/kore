package storage

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/billykore/kore/backend/internal/infra/storage/postgres"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	postgres.NewOrderRepository, wire.Bind(new(order.Repository), new(*postgres.OrderRepository)),
	postgres.NewOtpRepository, wire.Bind(new(otp.Repository), new(*postgres.OtpRepository)),
	postgres.NewProductRepository, wire.Bind(new(product.Repository), new(*postgres.ProductRepository)),
	postgres.NewShippingRepository, wire.Bind(new(shipping.Repository), new(*postgres.ShippingRepository)),
	postgres.NewUserRepository, wire.Bind(new(user.Repository), new(*postgres.UserRepository)),
)
