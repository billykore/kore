package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type OtpRepository interface {
	Get(ctx context.Context, otp model.Otp) (*model.Otp, error)
	Save(ctx context.Context, otp model.Otp) error
	Update(ctx context.Context, otp model.Otp) error
}
