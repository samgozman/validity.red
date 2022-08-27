package main

import (
	"os"
	"testing"

	"github.com/samgozman/validity.red/user/internal/models/user"
)

var testApp Config

func TestMain(m *testing.M) {
	repo := user.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
