package main

import (
	"os"
	"testing"

	document_mocks "github.com/samgozman/validity.red/document/mocks/models/document"
	notification_mocks "github.com/samgozman/validity.red/document/mocks/models/notification"
)

var testApp Config

func TestMain(m *testing.M) {
	testApp.Documents = document_mocks.NewDocumentDBTest(nil)
	testApp.Notifications = notification_mocks.NewNotificationDBTest(nil)
	os.Exit(m.Run())
}
