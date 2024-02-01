package usecase

import (
	"context"

	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/libs/repository"
)

type GreetUsecase struct {
	log  *log.Logger
	repo repository.Greeter
}

func NewGreetUsecase(log *log.Logger, repo repository.Greeter) *GreetUsecase {
	return &GreetUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *GreetUsecase) Greet(ctx context.Context, req *v1.GreetRequest) (*v1.GreetReply, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.GetName())
	return &v1.GreetReply{
		Message: "Hello " + req.GetName(),
	}, nil
}
