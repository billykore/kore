package usecase

import (
	"context"

	"github.com/billykore/kore/libs/entity"
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/repo"
)

type BookUsecase struct {
	log  *log.Logger
	repo repo.GreeterRepository
}

func NewBookUsecase(log *log.Logger, repo repo.GreeterRepository) *BookUsecase {
	return &BookUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *BookUsecase) Greet(ctx context.Context, req *entity.BookRequest) (*entity.BookResponse, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.Name)
	return &entity.BookResponse{
		Message: "Hello " + req.Name,
	}, nil
}
