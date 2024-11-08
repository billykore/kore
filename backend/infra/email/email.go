package email

import (
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/infra/email/mailer"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	mailer.NewOTPEmail, wire.Bind(new(otp.Email), new(*mailer.OTPEmail)),
)
