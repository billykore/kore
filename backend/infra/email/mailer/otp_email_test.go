package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOTPTemplate(t *testing.T) {
	tmpl, err := parseOTPTemplate("Oyen", "123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, tmpl)

	otpHtml := []byte(`<!DOCTYPE html>
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
    <p style="font-size:1.1em">Hi, Oyen!</p>
    <p>Please complete your login by enter the OTP. OTP is valid for 5 minutes</p>
    <h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">
        123456
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
</html>`)

	assert.Equal(t, otpHtml, tmpl)
}
