package mail

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSender_Send(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	assert.NoError(t, err)
	sender := &Mailer{
		from: os.Getenv("EMAIL_FROM"),
		host: os.Getenv("EMAIL_HOST"),
		port: port,
		key:  os.Getenv("EMAIL_PASSWORD"),
	}
	err = sender.Send(Data{
		Recipient: "billyimmcul2010@gmail.com",
		Subject:   "Test Email",
		Body: `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Verify Your Login</title>
</head>
<body>
<h1>Your OTP</h1>
<h3>123456</h3>
</body>
</html>`,
	})
	assert.NoError(t, err)
}
