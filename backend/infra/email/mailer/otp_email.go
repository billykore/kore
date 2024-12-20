package mailer

import (
	"bytes"
	"html/template"

	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/pkg/email/mailtrap"
	"github.com/billykore/kore/backend/pkg/logger"
)

var otpTmpl = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Verify Your Login</title>
</head>
<body style="max-width: 1024px;margin: 0 auto">
<div style="font-family: Helvetica,Arial,sans-serif;min-width:1000px;overflow:auto;line-height:2">
  <div style="margin:50px auto;width:70%;padding:20px 0">
    <div style="border-bottom:1px solid #eee">
      <a href="" style="font-size:1.4em;color: #00466a;text-decoration:none;font-weight:600">
        Kore Corp.
      </a>
    </div>
    <p style="font-size:1.1em">Hi, {{.Name}}!</p>
    <p>Please complete your login by enter the OTP. OTP is valid for 5 minutes</p>
    <h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">
        {{.OTP}}
    </h2>
    <p style="font-size:0.9em;">Regards,<br/>Kore Corp.</p>
    <hr style="border:none;border-top:1px solid #eee"/>
    <div style="float:right;padding:8px 0;color:#aaa;font-size:0.8em;line-height:1;font-weight:300">
      <p>Kore Corp.</p>
      <p>Jakarta</p>
      <p>Indonesia</p>
    </div>
  </div>
</div>
</body>
</html>`

// buffer to write the email template bytes.
var buffer = new(bytes.Buffer)

// parseOTPTemplate returns OTP html template.
func parseOTPTemplate(name, otp string) ([]byte, error) {
	defer buffer.Reset()
	tmpl := template.Must(template.New("otp").Parse(otpTmpl))
	err := tmpl.Execute(buffer, map[string]any{
		"Name": name,
		"OTP":  otp,
	})
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

type OTPEmail struct {
	log    *logger.Logger
	client *mailtrap.Client
}

func NewOTPEmail(log *logger.Logger, client *mailtrap.Client) *OTPEmail {
	return &OTPEmail{
		log:    log,
		client: client,
	}
}

func (e *OTPEmail) SendOTP(data otp.EmailData) error {
	body, err := parseOTPTemplate(data.Name, data.OTP)
	if err != nil {
		e.log.Usecase("SendOTP").Error(err)
		return err
	}
	err = e.client.Send(mailtrap.Data{
		Recipient: data.Recipient,
		Subject:   data.Subject,
		Body:      body,
	})
	if err != nil {
		e.log.Usecase("SendOTP").Error(err)
		return err
	}
	return nil
}
