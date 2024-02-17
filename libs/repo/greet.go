package repo

import (
	"context"

	"github.com/billykore/kore/libs/model"
)

type GreeterRepository interface {
	Get(ctx context.Context) (*model.Greet, error)
}
