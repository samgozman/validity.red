package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUserId         = status.Error(codes.InvalidArgument, "invalid user id")
	ErrInvalidDocumentId     = status.Error(codes.InvalidArgument, "invalid document_id")
	ErrInvalidNotificationId = status.Error(codes.InvalidArgument, "invalid notification_id")
	ErrDocumentNotFound      = status.Error(codes.NotFound, "document not found")
)
