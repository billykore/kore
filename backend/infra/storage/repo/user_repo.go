package repo

import (
	"context"
	"errors"
	"time"

	"github.com/billykore/kore/backend/domain/user"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

const duplicateKeyCode = "23505"

type UserRepo struct {
	postgres *gorm.DB
}

func NewUserRepo(postgres *gorm.DB) *UserRepo {
	return &UserRepo{postgres: postgres}
}

func (r *UserRepo) Create(ctx context.Context, u user.User) error {
	res := r.postgres.WithContext(ctx).Create(&u)
	if err := res.Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == duplicateKeyCode {
			return user.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	u := new(user.User)
	res := r.postgres.WithContext(ctx).
		Where("username = ?", username).
		First(u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) Login(ctx context.Context, auth user.AuthActivities) error {
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

func (r *UserRepo) SaveLoginActivity(ctx context.Context, auth user.AuthActivities) error {
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

func (r *UserRepo) Logout(ctx context.Context, auth user.AuthActivities) error {
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
		First(&auth)
	if err := res.Error; err != nil {
		return err
	}
	if auth.IsLoggedOut {
		return user.ErrAlreadyLoggedOut
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

func (r *UserRepo) FindFailedLoginByUsername(ctx context.Context, username string) (*user.AuthActivities, error) {
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
