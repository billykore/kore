package postgres

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewOrderRepository, wire.Bind(new(order.Repository), new(*OrderRepository)),
	NewOtpRepository, wire.Bind(new(otp.Repository), new(*OtpRepository)),
	NewProductRepository, wire.Bind(new(product.Repository), new(*ProductRepository)),
	NewShippingRepository, wire.Bind(new(shipping.Repository), new(*ShippingRepository)),
	NewUserRepository, wire.Bind(new(user.Repository), new(*UserRepository)),
)
