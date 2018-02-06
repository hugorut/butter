package mail

import (
	"gopkg.in/mailgun/mailgun-go.v1"
	"testing"
	"github.com/stretchr/testify/assert"
)

type args struct {
	from string
	subject string
	text string
	to []string
}

type mockMailSender struct {
		args []args
		message *mailgun.Message
}

func (m *mockMailSender) NewMessage(from, subject, text string, to ...string) *mailgun.Message {
		m.args = append(m.args, args{
			from: from,
			subject: subject,
			text: text,
			to: to,
		})

		return mailgun.NewMessage("d", "d", "t")
}

func (m *mockMailSender) Send(me *mailgun.Message) (string, string, error) {
	m.message = me

	return "", "", nil
}

func TestMailgunMailer_SetFrom_SendsWithNewScope(t *testing.T) {
	sender := &mockMailSender{}
	m := MailgunMailer{
		mg:   sender,
		from: "from-original",
	}
	x := m.SetFrom("from")
	x.Send("to", Message{}, 1)
	m.Send("to", Message{}, 1)

	assert.Equal(t, "from", sender.args[0].from)
	assert.Equal(t, "from-original", sender.args[1].from)
}



