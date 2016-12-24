package sys

import (
	"log"
)

const (
	INFO = iota
	WARN
	ERROR
	CRITICAL
	FATAL
)

var messages map[int]string = map[int]string{
	INFO:     "INFO",
	WARN:     "WARN",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
	FATAL:    "FATAL",
}

// Logger interface defines a entity that logs a message to a given output.
type Logger interface {
	Log(level int, message string)
}

type StdLogger struct {
	Levels map[int]string
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		Levels: messages,
	}
}

// StdLogger::Log logs message level to the default stdout.
func (s StdLogger) Log(level int, message string) {
	if _, ok := s.Levels[level]; !ok {
		level = INFO
	}

	if level == FATAL {
		log.Fatal(message)
	}

	log.Printf("[%s] ==> %s\n", s.Levels[level], message)
}
