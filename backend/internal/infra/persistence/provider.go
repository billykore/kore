package persistence

import (
	"github.com/billykore/kore/backend/internal/infra/persistence/postgres"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(postgres.New)
