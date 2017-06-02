package sys

import (
	"os"
	"testing"
)

func TestEnvOrDefault_ReturnsEnvValue(t *testing.T) {
	os.Setenv("test", "test-value")
	val := EnvOrDefault("test", "")

	if val != "test-value" {
		t.Errorf("failed asserting that key [test] had value 'test-value' instead was '%s'", val)
	}
}

func TestEnvOrDefault_ReturnsDefaultValueIfNoneSet(t *testing.T) {
	val := EnvOrDefault("test-not-set", "test-value")

	if val != "test-value" {
		t.Errorf("failed asserting that key [test-not-set] had value 'test-value' instead was '%s'", val)
	}
}
