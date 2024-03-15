package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type GreeterRepository interface {
	Get(ctx context.Context) (*model.Greet, error)
}
