package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/repo"
)

type ChatUsecase struct {
	log  *log.Logger
	repo repo.GreeterRepository
}

func NewChatUsecase(log *log.Logger, repo repo.GreeterRepository) *ChatUsecase {
	return &ChatUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ChatUsecase) Greet(ctx context.Context, req *entity.ChatRequest) (*entity.ChatResponse, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.Name)
	return &entity.ChatResponse{
		Message: "Hello " + req.Name,
	}, nil
}
