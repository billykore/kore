package db

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/billykore/kore/backend/internal/infra/db/postgres"
	"github.com/billykore/kore/backend/internal/infra/db/repository"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	postgres.New,
	repository.NewOrderRepository, wire.Bind(new(order.Repository), new(*repository.OrderRepository)),
	repository.NewOtpRepository, wire.Bind(new(otp.Repository), new(*repository.OtpRepository)),
	repository.NewProductRepository, wire.Bind(new(product.Repository), new(*repository.ProductRepository)),
	repository.NewShippingRepository, wire.Bind(new(shipping.Repository), new(*repository.ShippingRepository)),
	repository.NewUserRepository, wire.Bind(new(user.Repository), new(*repository.UserRepository)),
)
