package email

import (
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewOTPEmail, wire.Bind(new(otp.Email), new(*OTPEmail)),
	NewClient,
)
