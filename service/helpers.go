package service

import "os"

// get the system environment variable or fallback to a given value
func EnvOrDefault(name string, fallback string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}

	return fallback
}
