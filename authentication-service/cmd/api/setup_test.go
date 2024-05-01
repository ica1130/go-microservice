package main

import (
	"os"
	"testing"
)

// This is the entry point for the test suite
// It will run before any tests are run
func TestMain(m *testing.M) {

	// This will run the tests
	os.Exit(m.Run())
}
