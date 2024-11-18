package brevo

import (
	"github.com/billykore/kore/backend/pkg/config"
	"gopkg.in/gomail.v2"
)

// Data describe email data to send.
type Data struct {
	// Recipient is email recipient.
	Recipient string
	// Subject is subject of the email.
	Subject string
	// Body is the email body.
	Body []byte
}

// Client is Brevo email service client.
type Client struct {
	from string
	host string
	port int
	key  string
}

// NewClient returns new Client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		from: cfg.Email.From,
		host: cfg.Email.Host,
		port: cfg.Email.Port,
		key:  cfg.Email.Key,
	}
}

// Send sends email to the recipient.
func (c *Client) Send(data Data) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", c.from)
	msg.SetHeader("To", data.Recipient)
	msg.SetHeader("Subject", data.Subject)
	msg.SetBody("text/plain", string(data.Body))

	dialer := gomail.NewDialer(c.host, c.port, c.from, c.key)
	err := dialer.DialAndSend(msg)
	return err
}
