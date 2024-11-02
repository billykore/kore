package pkg

import (
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	logger.New,
	validation.New,
)
