package repo

import (
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(repo.ShippingRepository), new(*ShippingRepo)),
	NewShippingRepository,
)
