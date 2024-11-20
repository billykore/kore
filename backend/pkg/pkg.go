package pkg

import (
	"github.com/billykore/kore/backend/pkg/broker/rabbitmq"
	"github.com/billykore/kore/backend/pkg/db/postgres"
	"github.com/billykore/kore/backend/pkg/db/redis"
	"github.com/billykore/kore/backend/pkg/email/mailtrap"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	rabbitmq.NewConnection,
	postgres.New,
	mailtrap.NewClient,
	logger.New,
	validation.New,
	redis.New,
)
