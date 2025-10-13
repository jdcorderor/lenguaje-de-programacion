package testutil

import (
	"os"
	"testing"
)

// SuppressOutput suppresses stdout and stderr during tests
func SuppressOutput(t *testing.T) func() {
	stdoutRef := os.Stdout
	stderrRef := os.Stderr

	os.Stdout, _ = os.Open(os.DevNull)
	os.Stderr, _ = os.Open(os.DevNull)
	
	return func() {
		os.Stdout = stdoutRef
		os.Stderr = stderrRef
	}
}
