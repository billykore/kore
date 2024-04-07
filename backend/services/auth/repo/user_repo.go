package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type userRepo struct {
	postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) repo.UserRepository {
	return &userRepo{postgres: postgres}
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	res := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		First(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}
