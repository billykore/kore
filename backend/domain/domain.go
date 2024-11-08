package domain

import (
	"github.com/billykore/kore/backend/domain/order"
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/domain/product"
	"github.com/billykore/kore/backend/domain/shipping"
	"github.com/billykore/kore/backend/domain/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	order.NewService,
	otp.NewService,
	product.NewService,
	shipping.NewService,
	user.NewService,
)
