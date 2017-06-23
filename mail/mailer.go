package mail

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

// Mailer provides a simple interface for sending mail to an email
type Mailer interface {
	Send(to string, message Message, data interface{}) error
	SetFrom(string) Mailer
}

// Message holds the simple message to send
type Message struct {
	Subject string
	Html    string
	Plain   string
}

// ParseTemplate executes a template with the given data interface
// and then reads it into a bytes buffer
func ParseTemplate(text string, data interface{}) []byte {
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
