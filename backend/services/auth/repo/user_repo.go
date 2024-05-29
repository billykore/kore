package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) *UserRepo {
	return &UserRepo{postgres: postgres}
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	res := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		First(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}
