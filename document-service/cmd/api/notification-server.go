package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	"github.com/samgozman/validity.red/document/internal/utils"
	proto "github.com/samgozman/validity.red/document/proto"
)

type NotificationServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedNotificationServiceServer
}

func (ds *NotificationServer) Create(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()

	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), input.GetDocumentID())
	if err != nil {
		return nil, err
	}

	// create notification
	n := notification.Notification{
		DocumentID: documentID,
		Date:       input.GetDate().AsTime(),
	}
	err = ds.App.Notifications.InsertOne(ctx, &n)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' created notification '%s' successfully!", userID, n.ID)}
	return res, nil
}

func (ds *NotificationServer) Edit(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()

	// Validate input arguments
	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), input.GetDocumentID())
	if err != nil {
		return nil, err
	}
	notificationID, err := uuid.Parse(input.GetID())
	if err != nil {
		return nil, ErrInvalidNotificationId
	}

	// update notification
	n := notification.Notification{
		ID:         notificationID,
		DocumentID: documentID,
		Date:       input.GetDate().AsTime(),
	}
	err = ds.App.Notifications.UpdateOne(ctx, &n)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' edited notification '%s' successfully!", userID, n.ID)}
	return res, nil
}

func (ds *NotificationServer) Delete(ctx context.Context, req *proto.NotificationCreateRequest) (*proto.Response, error) {
	input := req.GetNotificationEntry()

	// Validate input arguments
	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), input.GetDocumentID())
	if err != nil {
		return nil, err
	}

	notificationID, err := uuid.Parse(input.GetID())
	if err != nil {
		return nil, ErrInvalidNotificationId
	}

	// delete notification
	n := notification.Notification{
		ID:         notificationID,
		DocumentID: documentID,
	}
	err = ds.App.Notifications.DeleteOne(ctx, &n)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' deleted notification with id '%s' successfully!", userID, n.ID)}
	return res, nil
}

// TODO: Refactor validators to reduce code duplication
func (ds *NotificationServer) GetAll(
	ctx context.Context,
	req *proto.NotificationsRequest,
) (*proto.ResponseNotificationsList, error) {
	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), req.GetDocumentID())
	if err != nil {
		return nil, err
	}

	// Find all notifications
	notifications, err := ds.App.Notifications.FindAll(ctx, documentID)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseNotificationsList{
		Result:        fmt.Sprintf("User '%s' found %d notifications successfully!", userID, len(notifications)),
		Notifications: utils.ConvertNotficationsToProtoFormat(&notifications),
	}
	return res, nil
}

func (ds *NotificationServer) Count(
	ctx context.Context,
	req *proto.NotificationsRequest,
) (*proto.ResponseCount, error) {
	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), req.GetDocumentID())
	if err != nil {
		return nil, err
	}

	count, err := ds.App.Notifications.Count(ctx, []uuid.UUID{documentID})
	if err != nil {
		return nil, err
	}

	res := &proto.ResponseCount{
		Result: fmt.Sprintf("User '%s' received notifications count for the '%s' document", userID, documentID),
		Count:  count,
	}
	return res, nil
}

// Helper to parse userId and documentId and validate document existence
func (ds *NotificationServer) checkInputsAndDocumentExistence(
	ctx context.Context,
	uID string,
	dID string,
) (
	userID uuid.UUID,
	documentID uuid.UUID,
	error error,
) {
	// Validate input arguments
	userID, err := uuid.Parse(uID)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrInvalidUserId
	}
	documentID, err = uuid.Parse(dID)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrInvalidDocumentId
	}

	// Check if that document exists
	d := document.Document{
		ID:     documentID,
		UserID: userID,
	}
	isDocumentExist, err := ds.App.Documents.Exists(ctx, &d)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	if !isDocumentExist {
		return uuid.Nil, uuid.Nil, ErrDocumentNotFound
	}

	return userID, documentID, nil
}
