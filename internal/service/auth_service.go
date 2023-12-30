package service

import (
	"context"

	v1 "github.com/billykore/todolist/internal/proto/v1"
	"github.com/billykore/todolist/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &v1.LoginReply{
		AccessToken: token.AccessToken,
		ExpiredTime: token.ExpiredTime,
	}, nil
}
