package domain

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/domain/product"
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	order.NewService,
	otp.NewService,
	product.NewService,
	shipping.NewService,
	user.NewService,
)
