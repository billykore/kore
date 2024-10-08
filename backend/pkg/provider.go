package pkg

import (
	"github.com/billykore/kore/backend/pkg/db"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/mail"
	"github.com/billykore/kore/backend/pkg/net/rabbit"
	"github.com/billykore/kore/backend/pkg/net/websocket"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewPostgres,
	log.NewLogger,
	websocket.NewPool,
	rabbit.New,
	mail.NewSender,
	validation.New,
)
