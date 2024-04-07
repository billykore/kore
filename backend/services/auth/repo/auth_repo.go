package repo

import (
	"context"
	"time"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/perrors"
	"github.com/billykore/kore/backend/pkg/repo"
	"gorm.io/gorm"
)

type authRepo struct {
	postgres *gorm.DB
}

func NewAuthRepository(postgres *gorm.DB) repo.AuthRepository {
	return &authRepo{postgres: postgres}
}

func (r *authRepo) Login(ctx context.Context, auth *model.AuthActivities) error {
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

func saveLoginActivity(ctx context.Context, tx *gorm.DB, auth *model.AuthActivities) error {
	res := tx.WithContext(ctx).Save(auth)
	err := res.Error
	return err
}

// autoLogout will logout all login activities.
func autoLogout(ctx context.Context, tx *gorm.DB, auth *model.AuthActivities) error {
	var authActivities []*model.AuthActivities
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

func (r *authRepo) Logout(ctx context.Context, auth *model.AuthActivities) error {
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

func saveLogoutActivity(ctx context.Context, tx *gorm.DB, auth *model.AuthActivities) error {
	res := tx.WithContext(ctx).
		Where("id = ?", auth.Id).
		Where("username = ?", auth.Username).
		First(auth)
	if err := res.Error; err != nil {
		return err
	}
	if auth.IsLoggedOut {
		return perrors.ErrAlreadyLoggedOut
	}
	return nil
}

func updateLogoutData(ctx context.Context, tx *gorm.DB, auth *model.AuthActivities) error {
	res := tx.WithContext(ctx).
		Model(auth).
		Where("id = ?", auth.Id).
		UpdateColumn("logout_time", time.Now()).
		UpdateColumn("is_logged_out", true)
	err := res.Error
	return err
}
