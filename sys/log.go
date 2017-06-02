package sys

import (
	"fmt"
	"log"
	"os"
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

// LogWriter is an interface to write a string to an output
type Writer interface {
	Printf(format string, v ...interface{})
	Fatal(v ...interface{})
	Print(v ...interface{})
}

// InMemWriter provides a lightweight writer for use with testing and bu
type InMemWriter struct {
	Log []string
}

// Fatal appends a formatted string to the memory log
func (m *InMemWriter) Printf(format string, v ...interface{}) {
	m.Log = append(m.Log, fmt.Sprintf(format, v...))
}

// Fatal appends the variadic interfaces to the memory log with a FATAL prefix
func (m *InMemWriter) Fatal(v ...interface{}) {
	slice := toStringSlice(v)

	for i := range slice {
		slice[i] = fmt.Sprintf("FATAL: %s", slice[i])
	}

	m.Log = append(m.Log, slice...)
}

// Print appends the variadic interfaces to the memory log
func (m *InMemWriter) Print(v ...interface{}) {
	m.Log = append(m.Log, toStringSlice(v)...)
}

// toStringSlice converts a variadic list of interfaces into a slice of strings
func toStringSlice(vals ...interface{}) []string {
	var output []string

	for _, v := range vals {
		output = append(output, fmt.Sprintf("%v", v))
	}

	return output
}

// StdLogger wraps the stdout in the logger interface
type StdLogger struct {
	levels map[int]string
	writer Writer
}

// NewStdLogger returns a new instance of the the StdLogger with the correct levels set
func NewStdLogger() *StdLogger {
	var std = log.New(os.Stderr, "", log.LstdFlags)
	return &StdLogger{
		levels: messages,
		writer: std,
	}
}

// Log logs message level to the default stdout.
func (s StdLogger) Log(level int, message string) {
	if _, ok := s.levels[level]; !ok {
		level = INFO
	}

	if level == FATAL {
		s.writer.Fatal(message)
	}

	s.writer.Printf("[%s] ==> %s\n", s.levels[level], message)
}

// TestLogger provides a stub to ease testing
type TestLogger struct {
	messages []struct {
		level int
		m     string
	}
}

// Log appends a message to the underlying messages property
func (t *TestLogger) Log(level int, message string) {
	t.messages = append(t.messages, struct {
		level int
		m     string
	}{
		level: level,
		m:     message,
	})
}

//AssertCalled asserts that a log level and message combination was called
func (t *TestLogger) AssertCalled(level int, message string) bool {
	for _, m := range t.messages {
		if m.level == level && m.m == message {
			return true
		}
	}

	return false
}
