package mail

import mailgun "gopkg.in/mailgun/mailgun-go.v1"

// MailgunMailer is a struct that wraps the Mailgun api in the
// Mailer interface
type MailgunMailer struct {
	mg   MailgunMessageSender
	from string
}

// MailgunMessageSender specifies an interface for sending mailgun messages
type MailgunMessageSender interface {
	NewMessage(from, subject, text string, to ...string) *mailgun.Message
	Send(m *mailgun.Message) (string, string, error)
}

// SetFrom sets the value of the mailer from
func (m *MailgunMailer) SetFrom(from string) Mailer {
	return &MailgunMailer{
		mg:m.mg,
		from: from,
	}
}

// Send provides a method to quickly send a message via the mailgun API
func (m *MailgunMailer) Send(to string, message Message, data interface{}) error {
	plainBody := ParseTemplate(message.Plain, data)

	mail := m.mg.NewMessage(m.from, message.Subject, string(plainBody), to)

	body := ParseTemplate(message.Html, data)
	mail.SetHtml(string(body))

	_, _, err := m.mg.Send(mail)
	return err
}
