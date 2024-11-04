package pkg

import (
	"github.com/billykore/kore/backend/pkg/broker/rabbitmq"
	"github.com/billykore/kore/backend/pkg/db/postgres"
	"github.com/billykore/kore/backend/pkg/email/brevo"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	rabbitmq.NewConnection,
	postgres.New,
	brevo.NewClient,
	logger.New,
	validation.New,
)
