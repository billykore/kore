package pkg

import (
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/websocket"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewPostgres,
	log.NewLogger,
	websocket.NewPool,
)
