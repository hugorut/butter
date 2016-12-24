package mail

import (
	"os"

	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

// NewMailer returns the mailer set up in the configuration files
// the default mailer is the mailgun implementation
func NewMailer() Mailer {
	switch os.Getenv("MAIL_PROVIDER") {
	case "mailgun":
		return NewMailgunMailer()
	}

	return NewMailgunMailer()
}

// NewMailgunMailer returns new MailgunMailer from environment variables
func NewMailgunMailer() Mailer {
	return &MailgunMailer{
		mailgun.NewMailgun(
			os.Getenv("MAILGUN_DOMAIN"),
			os.Getenv("MAILGUN_API_KEY"),
			os.Getenv("MAILGUN_PUB_KEY"),
		),
		os.Getenv("FROM_EMAIL"),
	}
}
