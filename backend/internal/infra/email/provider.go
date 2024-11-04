package email

import (
	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/infra/email/brevo"
	"github.com/billykore/kore/backend/internal/infra/email/mailer"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	mailer.NewOTPEmail, wire.Bind(new(otp.Email), new(*mailer.OTPEmail)),
	brevo.NewClient,
)
