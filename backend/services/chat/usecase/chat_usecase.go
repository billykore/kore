package usecase

import (
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
