package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/billykore/kore/backend/pkg/pkgerr"
	"gorm.io/gorm"
)

type UserRepository struct {
	postgres *gorm.DB
}

func NewUserRepository(postgres *gorm.DB) *UserRepository {
	return &UserRepository{postgres: postgres}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	u := new(user.User)
	res := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		First(u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Login(ctx context.Context, auth user.AuthActivities) error {
	tx := r.postgres.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	err := autoLogout(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = saveLoginActivity(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *UserRepository) SaveLoginActivity(ctx context.Context, auth user.AuthActivities) error {
	tx := r.postgres.Begin()
	err := saveLoginActivity(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func saveLoginActivity(ctx context.Context, tx *gorm.DB, auth user.AuthActivities) error {
	res := tx.WithContext(ctx).Save(&auth)
	err := res.Error
	return err
}

// autoLogout will logout all login activities.
func autoLogout(ctx context.Context, tx *gorm.DB, auth user.AuthActivities) error {
	var authActivities []user.AuthActivities
	res := tx.WithContext(ctx).
		Where("username = ?", auth.Username).
		Find(&authActivities)
	if err := res.Error; err != nil {
		return err
	}
	if authActivities != nil && len(authActivities) > 0 {
		for _, a := range authActivities {
			err := updateLogoutData(ctx, tx, a)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *UserRepository) Logout(ctx context.Context, auth user.AuthActivities) error {
	tx := r.postgres.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	err := saveLogoutActivity(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = updateLogoutData(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func saveLogoutActivity(ctx context.Context, tx *gorm.DB, auth user.AuthActivities) error {
	res := tx.WithContext(ctx).
		Where("uuid = ?", auth.UUID).
		Where("username = ?", auth.Username).
		First(auth)
	if err := res.Error; err != nil {
		return err
	}
	if auth.IsLoggedOut {
		return pkgerr.ErrAlreadyLoggedOut
	}
	return nil
}

func updateLogoutData(ctx context.Context, tx *gorm.DB, auth user.AuthActivities) error {
	res := tx.WithContext(ctx).
		Model(auth).
		Where("uuid = ?", auth.UUID).
		UpdateColumn("logout_time", time.Now()).
		UpdateColumn("is_logged_out", true)
	err := res.Error
	return err
}

func (r *UserRepository) FindFailedLoginByUsername(ctx context.Context, username string) (*user.AuthActivities, error) {
	activities := new(user.AuthActivities)
	tx := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		Where("is_login_succeed = false").
		First(&activities)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return activities, tx.Error
}
