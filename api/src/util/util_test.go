package util

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	// Save environment to put it back
	prevPort := os.Getenv(portEnv)
	// Clear port to check default
	os.Setenv(portEnv, "")
	if GetPort() != 8080 {
		t.Error("Default port was not returned with no env var set")
	}
	os.Setenv(portEnv, "9999")
	if GetPort() != 9999 {
		t.Error("Port environment variable was ignored")
	}
	os.Setenv(portEnv, "NaN")
	if GetPort() != 8080 {
		t.Error("Invalid port was not ignored")
	}
	// Set port back
	os.Setenv(portEnv, prevPort)
}
