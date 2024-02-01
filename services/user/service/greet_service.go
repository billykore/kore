package service

import (
	"context"

	"github.com/billykore/kore/services/user/usecase"
	"github.com/billykore/kore/libs/proto/v1"
)

type UserService struct {
	v1.UnimplementedGreeterServer

	uc *usecase.GreetUsecase
}

func NewUserService(uc *usecase.GreetUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Greet(ctx context.Context, in *v1.GreetRequest) (*v1.GreetReply, error) {
	greet, err := s.uc.Greet(ctx, in)
	if err != nil {
		return nil, err
	}
	return greet, nil
}
