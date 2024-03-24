package pkg

import (
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.New,
	log.NewLogger,
)
