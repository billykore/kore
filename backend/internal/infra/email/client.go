package email

import (
	"github.com/billykore/kore/backend/pkg/config"
	"gopkg.in/gomail.v2"
)

type Data struct {
	Recipient string
	Subject   string
	Body      []byte
}

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

// Send email to recipient.
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
