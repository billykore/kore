package repo

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/domain/otp"
	"gorm.io/gorm"
)

type OtpRepo struct {
	db *gorm.DB
}

func NewOtpRepo(db *gorm.DB) *OtpRepo {
	return &OtpRepo{
		db: db,
	}
}

func (r *OtpRepo) Get(ctx context.Context, otp otp.OTP) (*otp.OTP, error) {
	return r.getByEmailAndValue(ctx, otp.Email, otp.Otp)
}

func (r *OtpRepo) getByEmailAndValue(ctx context.Context, email, value string) (*otp.OTP, error) {
	o := new(otp.OTP)
	tx := r.db.WithContext(ctx).
		Where("email = ?", email).
		Where("otp = ?", value).
		First(o)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return o, nil
}

func (r *OtpRepo) Save(ctx context.Context, otp otp.OTP) error {
	tx := r.db.WithContext(ctx).Save(&otp)
	return tx.Error
}

func (r *OtpRepo) Update(ctx context.Context, otp otp.OTP) error {
	return r.deactivateOtp(ctx, otp.Otp)
}

func (r *OtpRepo) deactivateOtp(ctx context.Context, otpValue string) error {
	o := new(otp.OTP)
	tx := r.db.WithContext(ctx).
		Model(o).
		Where("otp = ?", otpValue).
		Update("is_active", false)
	return tx.Error
}
