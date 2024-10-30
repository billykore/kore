package otp

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, otp Otp) (*Otp, error)
	Save(ctx context.Context, otp Otp) error
	Update(ctx context.Context, otp Otp) error
}
