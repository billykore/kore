package usecase

import (
	"github.com/billykore/kore/pkg/log"
	"github.com/billykore/kore/pkg/repo"
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
