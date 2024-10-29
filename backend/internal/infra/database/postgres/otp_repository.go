package postgres

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/internal/domain/otp"
	"gorm.io/gorm"
)

type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) *OtpRepository {
	return &OtpRepository{
		db: db,
	}
}

func (r *OtpRepository) Get(ctx context.Context, otp otp.Otp) (*otp.Otp, error) {
	return r.getByEmailAndValue(ctx, otp.Email, otp.Otp)
}

func (r *OtpRepository) getByEmailAndValue(ctx context.Context, email, value string) (*otp.Otp, error) {
	o := new(otp.Otp)
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

func (r *OtpRepository) Save(ctx context.Context, otp otp.Otp) error {
	tx := r.db.WithContext(ctx).Save(&otp)
	return tx.Error
}

func (r *OtpRepository) Update(ctx context.Context, otp otp.Otp) error {
	return r.deactivateOtp(ctx, otp.Otp)
}

func (r *OtpRepository) deactivateOtp(ctx context.Context, otpValue string) error {
	o := new(otp.Otp)
	tx := r.db.WithContext(ctx).
		Model(o).
		Where("otp = ?", otpValue).
		Update("is_active", false)
	return tx.Error
}
