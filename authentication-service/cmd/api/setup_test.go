package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp Config

// This is the entry point for the test suite
// It will run before any tests are run
func TestMain(m *testing.M) {

	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo

	// This will run the tests
	os.Exit(m.Run())
}
