package repository

import (
	"context"

	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/repository"
)

type greetRepo struct {
}

func NewGreetRepository() repository.Greeter {
	return &greetRepo{}
}

func (r *greetRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
