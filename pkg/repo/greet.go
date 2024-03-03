package repo

import (
	"context"

	"github.com/billykore/kore/pkg/model"
)

type GreeterRepository interface {
	Get(ctx context.Context) (*model.Greet, error)
}
