package mail

import (
	"github.com/billykore/kore/backend/pkg/config"
	"gopkg.in/gomail.v2"
)

type TemplateFunc func(any) ([]byte, error)

type Data struct {
	Recipient string
	Subject   string
	Body      []byte
}

type Mailer struct {
	from string
	host string
	port int
	key  string
}

func NewSender(cfg *config.Config) *Mailer {
	return &Mailer{
		from: cfg.Email.From,
		host: cfg.Email.Host,
		port: cfg.Email.Port,
		key:  cfg.Email.Key,
	}
}

// Send mail to user.
func (s *Mailer) Send(data Data) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", data.Recipient)
	msg.SetHeader("Subject", data.Subject)
	msg.SetBody("text/plain", string(data.Body))

	dialer := gomail.NewDialer(s.host, s.port, s.from, s.key)
	err := dialer.DialAndSend(msg)
	return err
}
