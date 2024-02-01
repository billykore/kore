package tpl

func UsecaseProviderTemplate() []byte {
	return []byte(`package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGreetUsecase)
`)
}

func UsecaseTemplate() []byte {
	return []byte(`package usecase

import (
	"context"

	"{{ .Mod }}/libs/pkg/log"
	"{{ .Mod }}/libs/proto/v1"
	"{{ .Mod }}/libs/repository"
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
`)
}
