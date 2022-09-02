package main

import "errors"

var (
	ErrInvalidUserId         = errors.New("invalid user id")
	ErrInvalidDocumentId     = errors.New("invalid document_id")
	ErrInvalidNotificationId = errors.New("invalid notification_id")
	ErrDocumentNotFound      = errors.New("document does not exist")
)
