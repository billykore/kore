package user

import (
	"context"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
	Login(ctx context.Context, auth AuthActivities) error
	SaveLoginActivity(ctx context.Context, auth AuthActivities) error
	Logout(ctx context.Context, auth AuthActivities) error
	FindFailedLoginByUsername(ctx context.Context, username string) (*AuthActivities, error)
}
