package mail

import mailgun "gopkg.in/mailgun/mailgun-go.v1"

// MailgunMailer is a struct that wraps the Mailgun api in the
// Mailer interface
type MailgunMailer struct {
	mg   mailgun.Mailgun
	from string
}

// SetFrom sets the value of the mailer from
func (m *MailgunMailer) SetFrom(from string) Mailer {
	m.from = from

	return m
}

// Send provides a method to quickly send a message via the mailgun API
func (m *MailgunMailer) Send(to string, message Message, data interface{}) error {
	plainBody := ParseTemplate(message.Plain, data)
	mailgunMessage := mailgun.NewMessage(m.from, message.Subject, string(plainBody), to)

	body := ParseTemplate(message.Html, data)
	mailgunMessage.SetHtml(string(body))

	_, _, err := m.mg.Send(mailgunMessage)
	return err
}
