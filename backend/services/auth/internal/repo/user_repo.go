package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) *UserRepository {
	return &UserRepository{postgres: postgres}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	res := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		First(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}
