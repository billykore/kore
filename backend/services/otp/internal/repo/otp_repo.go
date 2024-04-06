package repo

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/pkg/model"
	"gorm.io/gorm"
)

type OtpRepo struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) *OtpRepo {
	return &OtpRepo{
		db: db,
	}
}

func (r *OtpRepo) Get(ctx context.Context, otp model.Otp) (*model.Otp, error) {
	return r.getByEmailAndValue(ctx, otp.Email, otp.Otp)
}

func (r *OtpRepo) getByEmailAndValue(ctx context.Context, email, value string) (*model.Otp, error) {
	otp := new(model.Otp)
	tx := r.db.WithContext(ctx).
		Where("email = ?", email).
		Where("otp = ?", value).
		First(otp)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return otp, nil
}

func (r *OtpRepo) Save(ctx context.Context, otp model.Otp) error {
	tx := r.db.WithContext(ctx).Save(&otp)
	return tx.Error
}

func (r *OtpRepo) Update(ctx context.Context, otp model.Otp) error {
	return r.deactivateOtp(ctx, otp.Otp)
}

func (r *OtpRepo) deactivateOtp(ctx context.Context, otpValue string) error {
	otp := new(model.Otp)
	tx := r.db.WithContext(ctx).
		Model(otp).
		Where("otp = ?", otpValue).
		Update("is_active", false)
	return tx.Error
}
