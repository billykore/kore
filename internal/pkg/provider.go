package pkg

import (
	"github.com/billykore/todolist/internal/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(log.NewLogger)
