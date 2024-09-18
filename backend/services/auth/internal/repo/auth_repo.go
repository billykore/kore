package repo

import (
	"context"
	"errors"
	"time"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/svcerr"
	"gorm.io/gorm"
)

type AuthRepository struct {
	postgres *gorm.DB
}

func NewAuthRepository(postgres *gorm.DB) *AuthRepository {
	return &AuthRepository{postgres: postgres}
}

func (r *AuthRepository) Login(ctx context.Context, auth model.AuthActivities) error {
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

func (r *AuthRepository) SaveLoginActivity(ctx context.Context, auth model.AuthActivities) error {
	tx := r.postgres.Begin()
	err := saveLoginActivity(ctx, tx, auth)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func saveLoginActivity(ctx context.Context, tx *gorm.DB, auth model.AuthActivities) error {
	res := tx.WithContext(ctx).Save(&auth)
	err := res.Error
	return err
}

// autoLogout will logout all login activities.
func autoLogout(ctx context.Context, tx *gorm.DB, auth model.AuthActivities) error {
	var authActivities []model.AuthActivities
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

func (r *AuthRepository) Logout(ctx context.Context, auth model.AuthActivities) error {
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

func saveLogoutActivity(ctx context.Context, tx *gorm.DB, auth model.AuthActivities) error {
	res := tx.WithContext(ctx).
		Where("uuid = ?", auth.UUID).
		Where("username = ?", auth.Username).
		First(auth)
	if err := res.Error; err != nil {
		return err
	}
	if auth.IsLoggedOut {
		return svcerr.ErrAlreadyLoggedOut
	}
	return nil
}

func updateLogoutData(ctx context.Context, tx *gorm.DB, auth model.AuthActivities) error {
	res := tx.WithContext(ctx).
		Model(auth).
		Where("uuid = ?", auth.UUID).
		UpdateColumn("logout_time", time.Now()).
		UpdateColumn("is_logged_out", true)
	err := res.Error
	return err
}

func (r *AuthRepository) FindFailedLoginByUsername(ctx context.Context, username string) (*model.AuthActivities, error) {
	activities := new(model.AuthActivities)
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
