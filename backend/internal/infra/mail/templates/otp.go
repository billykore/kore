package templates

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed otp.gohtml
var otpTmpl embed.FS

// buffer to write the email template bytes.
var buffer = new(bytes.Buffer)

// OtpData contains data for OTP html template.
type OtpData struct {
	Otp string
}

// OtpTemplate returns OTP html template with OtpData.
func OtpTemplate(data OtpData) ([]byte, error) {
	tmpl, err := template.ParseFS(otpTmpl, "otp.gohtml")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
