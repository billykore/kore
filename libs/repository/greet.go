package repository

import (
	"context"

	"github.com/billykore/kore/libs/model"
)

type Greeter interface {
	Get(ctx context.Context) (*model.Greet, error)
}
