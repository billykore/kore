package pkg

import (
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(log.NewLogger)
