package sys

import "testing"

func TestStdLogger_Log_WritesWithGivenLevel(t *testing.T) {
	type test struct {
		inputLevel   int
		inputMessage string
		expectedLog  string
	}

	tests := []test{
		{
			inputLevel:   INFO,
			inputMessage: "my message",
			expectedLog:  "[INFO] ==> my message\n",
		},
		{
			inputLevel:   WARN,
			inputMessage: "my message",
			expectedLog:  "[WARN] ==> my message\n",
		},
		{
			inputLevel:   ERROR,
			inputMessage: "my message",
			expectedLog:  "[ERROR] ==> my message\n",
		},
		{
			inputLevel:   CRITICAL,
			inputMessage: "my message",
			expectedLog:  "[CRITICAL] ==> my message\n",
		},
	}

	for _, ts := range tests {
		memWriter := &InMemWriter{}

		logger := StdLogger{
			levels: messages,
			writer: memWriter,
		}

		logger.Log(ts.inputLevel, ts.inputMessage)

		if memWriter.Log[0] != ts.expectedLog {
			t.Errorf("Logger did not write correct error message to log\nExpected: %s\nGot:%s", ts.expectedLog, memWriter.Log[0])
		}
	}
}

func TestStdLogger_Log_writesFatal_ifLevelFatal(t *testing.T) {
	memWriter := &InMemWriter{}

	logger := StdLogger{
		levels: messages,
		writer: memWriter,
	}

	logger.Log(FATAL, "my fatal message")
	expectedMessage := "FATAL: [my fatal message]"

	if memWriter.Log[0] != expectedMessage {
		t.Errorf("Logger did not write correct error message to log\nExpected: %s\nGot:%s", expectedMessage, memWriter.Log[0])
	}
}
