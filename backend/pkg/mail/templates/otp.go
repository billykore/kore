package templates

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed otp.gohtml
var otpTmpl embed.FS

// buffer to write the email template
var buffer = new(bytes.Buffer)

// OtpData contains data for OTP html template.
type OtpData struct {
	Otp string
}

// OtpTemplate returns OTP html template with OtpData.
func OtpTemplate(data OtpData) (string, error) {
	tmpl, err := template.ParseFS(otpTmpl, "otp.gohtml")
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
