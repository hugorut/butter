package sys

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

type Mailer interface {
	Send(to string, subject string, html string, plain string, data interface{}) error
}

type MailgunMailer struct {
	mg   mailgun.Mailgun
	from string
}

func NewMailer() Mailer {
	return MailgunMailer{
		mailgun.NewMailgun(
			os.Getenv("MAILGUN_DOMAIN"),
			os.Getenv("MAILGUN_API_KEY"),
			os.Getenv("MAILGUN_PUB_KEY"),
		),
		os.Getenv("FROM_EMAIL"),
	}
}

// Send provides a method to quickly send a message via the mailgun API
func (m *MailgunMailer) Send(to string, subject string, html string, plain string, data interface{}) error {
	plainBody := parseTemplate(plain, data)
	message := mailgun.NewMessage(m.from, subject, string(plainBody), to)

	body := parseTemplate(html, data)
	message.SetHtml(string(body))

	_, _, err := m.mg.Send(message)
	return err
}

func parseTemplate(text string, data interface{}) []byte {
	output := new(bytes.Buffer)

	t, err := template.New("email").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(output, data)

	body, err := ioutil.ReadAll(output)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
