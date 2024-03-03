package pkg

import (
	"github.com/billykore/kore/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(log.NewLogger)
