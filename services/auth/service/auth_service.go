package service

import (
	"context"

	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/services/auth/usecase"
)

type AuthService struct {
	v1.UnimplementedAuthorizationServer

	uc *usecase.AuthUsecase
}

func NewAuthService(uc *usecase.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := s.uc.Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}
