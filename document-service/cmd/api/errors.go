package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUserID         = status.Error(codes.InvalidArgument, "invalid user id")
	ErrInvalidDocumentID     = status.Error(codes.InvalidArgument, "invalid document_id")
	ErrInvalidNotificationID = status.Error(codes.InvalidArgument, "invalid notification_id")
	ErrDocumentNotFound      = status.Error(codes.NotFound, "document not found")
	ErrMaxDocumentsLimit     = status.Error(codes.Canceled, "max documents limit reached")
	ErrMaxNotificationsLimit = status.Error(codes.Canceled, "max notifications for this document limit reached")
)
