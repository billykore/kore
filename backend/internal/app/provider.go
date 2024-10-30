package app

import (
	"github.com/billykore/kore/backend/internal/app/order"
	"github.com/billykore/kore/backend/internal/app/otp"
	"github.com/billykore/kore/backend/internal/app/product"
	"github.com/billykore/kore/backend/internal/app/shipping"
	"github.com/billykore/kore/backend/internal/app/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	order.NewService,
	otp.NewService,
	product.NewService,
	shipping.NewService,
	user.NewService,
)
