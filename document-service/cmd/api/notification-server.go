package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"gorm.io/gorm"
)

type NotificationServer struct {
	db *gorm.DB
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedNotificationServiceServer
}

func (ds *NotificationServer) Create(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()
	inputId := req.GetUserID()

	userID, err := uuid.Parse(inputId)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	documentID, err := uuid.Parse(input.GetDocumentID())
	if err != nil {
		return nil, errors.New("invalid document_id")
	}

	// Check if that document exists
	d := document.Document{
		ID:     documentID,
		UserID: userID,
	}
	isDocumentExist, err := d.Exists(ctx, ds.db)
	if err != nil {
		return nil, err
	}
	if !isDocumentExist {
		return nil, errors.New("document does not exist")
	}

	// create notification
	n := notification.Notification{
		DocumentID: documentID,
		Date:       input.GetDate().AsTime(),
	}
	err = n.InsertOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Notification with id '%s' created successfully!", n.ID)}
	return res, nil
}

func (ds *NotificationServer) Edit(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()

	// Validate input arguments
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	documentID, err := uuid.Parse(input.GetDocumentID())
	if err != nil {
		return nil, errors.New("invalid document_id")
	}
	notificationID, err := uuid.Parse(input.GetID())
	if err != nil {
		return nil, errors.New("invalid notification id")
	}

	// Check if that document exists
	d := document.Document{
		ID:     documentID,
		UserID: userID,
	}
	isDocumentExist, err := d.Exists(ctx, ds.db)
	if err != nil {
		return nil, err
	}
	if !isDocumentExist {
		return nil, errors.New("document does not exist")
	}

	// update notification
	n := notification.Notification{
		ID:         notificationID,
		DocumentID: documentID,
		Date:       input.GetDate().AsTime(),
	}
	err = n.UpdateOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Notification with id '%s' updated successfully!", n.ID)}
	return res, nil
}

func (ds *NotificationServer) Delete(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()

	// Validate input arguments
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	documentID, err := uuid.Parse(input.GetDocumentID())
	if err != nil {
		return nil, errors.New("invalid document_id")
	}
	notificationID, err := uuid.Parse(input.GetID())
	if err != nil {
		return nil, errors.New("invalid notification id")
	}

	// Check if that document exists
	d := document.Document{
		ID:     documentID,
		UserID: userID,
	}
	isDocumentExist, err := d.Exists(ctx, ds.db)
	if err != nil {
		return nil, err
	}
	if !isDocumentExist {
		return nil, errors.New("document does not exist")
	}

	// delete notification
	n := notification.Notification{
		ID:         notificationID,
		DocumentID: documentID,
	}
	err = n.DeleteOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Notification with id '%s' deleted successfully!", n.ID)}
	return res, nil
}
