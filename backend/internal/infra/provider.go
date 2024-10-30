package infra

import (
	"github.com/billykore/kore/backend/internal/infra/http"
	"github.com/billykore/kore/backend/internal/infra/mail"
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbit"
	"github.com/billykore/kore/backend/internal/infra/persistence"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	persistence.NewPostgres,
	mail.NewSender,
	logger.New,
	http.NewServer,
	http.NewRouter,
	rabbit.New,
)
