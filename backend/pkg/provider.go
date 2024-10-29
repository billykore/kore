package pkg

import (
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	validation.New,
)
