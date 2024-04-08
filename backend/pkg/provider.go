package pkg

import (
	"github.com/billykore/kore/pkg/db"
	"github.com/billykore/kore/pkg/log"
	"github.com/billykore/kore/pkg/websocket"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewPostgres,
	log.NewLogger,
	websocket.NewPool,
)
