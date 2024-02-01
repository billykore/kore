package pkg

import (
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(log.NewLogger)
