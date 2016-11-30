package sys

import (
	"log"
)

// Logger interface defines a entity that logs a message to a given output.
type Logger interface {
	Log(level, message string)
}

type StdLogger struct{}

// StdLogger::Log logs message level to the default stdout.
func (s StdLogger) Log(level, message string) {
	if level == "fatal" {
		log.Fatal(message)
	}

	log.Printf("%s ==> %s \n", level, message)
}
