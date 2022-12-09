package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	"github.com/samgozman/validity.red/document/internal/utils"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedNotificationServiceServer
}

func (ds *NotificationServer) Create(ctx context.Context, req *proto.NotificationCreateRequest) (*emptypb.Empty, error) {
	input := req.GetNotificationEntry()

	userID, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), input.GetDocumentID())
	if err != nil {
		return nil, err
	}

	// Check if user has reached the limit of notifications
	count, err := ds.App.Notifications.Count(ctx, documentID)
	if err != nil {
		return nil, err
	}
	if count >= ds.App.limits.MaxNotificationsPerDocument {
		return nil, ErrMaxNotificationsLimit
	}
	// create notification
	n := notification.Notification{
		UserID:     userID,
		DocumentID: documentID,
		Date:       input.GetDate().AsTime(),
	}
	err = ds.App.Notifications.InsertOne(ctx, &n)

	// return error if exists
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ds *NotificationServer) Delete(ctx context.Context, req *proto.NotificationCreateRequest) (*emptypb.Empty, error) {
	input := req.GetNotificationEntry()

	// Validate input arguments
	_, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), input.GetDocumentID())
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

	return &emptypb.Empty{}, nil
}

func (ds *NotificationServer) GetAll(
	ctx context.Context,
	req *proto.NotificationsRequest,
) (*proto.ResponseNotificationsList, error) {
	_, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), req.GetDocumentID())
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
		Notifications: utils.ConvertNotificationsToProtoFormat(&notifications),
	}
	return res, nil
}

func (ds *NotificationServer) Count(
	ctx context.Context,
	req *proto.NotificationsCountRequest,
) (*proto.ResponseCount, error) {
	_, documentID, err := ds.checkInputsAndDocumentExistence(ctx, req.GetUserID(), req.GetDocumentID())
	if err != nil {
		return nil, err
	}

	count, err := ds.App.Notifications.Count(ctx, documentID)
	if err != nil {
		return nil, err
	}

	res := &proto.ResponseCount{
		Count: count,
	}
	return res, nil
}

func (ds *NotificationServer) CountAll(
	ctx context.Context,
	req *proto.NotificationsAllRequest,
) (*proto.ResponseCount, error) {
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserId
	}

	count, err := ds.App.Notifications.CountAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := &proto.ResponseCount{
		Count: count,
	}
	return res, nil
}

func (ds *NotificationServer) GetAllForUser(
	ctx context.Context,
	req *proto.NotificationsAllRequest,
) (*proto.ResponseNotificationsList, error) {
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserId
	}

	// Find all notifications
	notifications, err := ds.App.Notifications.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseNotificationsList{
		Notifications: utils.ConvertNotificationsToProtoFormat(&notifications),
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
