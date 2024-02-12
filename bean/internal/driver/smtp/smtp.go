package smtp

import (
	"fmt"
	"net/smtp"

	"harvest/bean/internal/usecase/interfaces"
)

type client struct {
	addr string
	auth smtp.Auth
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

func New(cfg *Config) interfaces.Emailer {
	return &client{
		addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		auth: smtp.CRAMMD5Auth(cfg.Username, cfg.Password),
	}
}

func (c *client) Send(
	from string,
	to string,
	subject string,
	body string,
) error {
	msg := fmt.Sprintf(
		("From: %s" +
			"\r\n" +
			"To: %s" +
			"\r\n" +
			"Subject: %s" +
			"\r\n" +
			"\r\n" +
			"%s" +
			"\r\n"),
		from, to, subject, body,
	)

	err := smtp.SendMail(c.addr, c.auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
