package main

import (
	"os"
	"testing"

	"github.com/samgozman/validity.red/user/mocks"
)

var testApp Config

func TestMain(m *testing.M) {
	repo := mocks.NewPostgresTestRepository(nil)
	testApp.Repo = repo

	os.Exit(m.Run())
}
