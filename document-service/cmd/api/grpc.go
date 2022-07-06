package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	"github.com/samgozman/validity.red/document/internal/utils"
	proto "github.com/samgozman/validity.red/document/proto"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DocumentServer struct {
	db *gorm.DB
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedDocumentServiceServer
}

type NotificationServer struct {
	db *gorm.DB
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedNotificationServiceServer
}

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterDocumentServiceServer(s, &DocumentServer{
		db: app.db,
	})
	proto.RegisterNotificationServiceServer(s, &NotificationServer{
		db: app.db,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (ds *DocumentServer) Create(ctx context.Context, req *proto.DocumentCreateRequest) (*proto.Response, error) {
	input := req.GetDocumentEntry()

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// register document
	d := document.Document{
		UserID:      userID,
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	}
	err = d.InsertOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Document with title '%s' created successfully!", input.Title)}
	return res, nil
}

func (ds *DocumentServer) Edit(ctx context.Context, req *proto.DocumentCreateRequest) (*proto.Response, error) {
	input := req.GetDocumentEntry()

	// Decode values
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, errors.New("invalid document id")
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// update document
	d := document.Document{
		ID:          id,
		UserID:      userID,
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	}
	err = d.UpdateOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Document with title '%s' updated successfully!", input.Title)}
	return res, nil
}

func (ds *DocumentServer) Delete(ctx context.Context, req *proto.DocumentRequest) (*proto.Response, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, errors.New("invalid document id")
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// delete document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = d.DeleteOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Document with id '%s' deleted successfully!", id)}
	return res, nil
}

func (ds *DocumentServer) GetOne(ctx context.Context, req *proto.DocumentRequest) (*proto.ResponseDocument, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, errors.New("invalid document id")
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// Find document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = d.FindOne(ctx, ds.db)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocument{
		Result: fmt.Sprintf("Found Document with title '%s' successfully!", d.Title),
		Document: &proto.Document{
			ID:          d.ID.String(),
			UserID:      d.UserID.String(),
			Title:       d.Title,
			Type:        d.Type,
			Description: d.Description,
			ExpiresAt:   timestamppb.New(d.ExpiresAt),
		},
		Notifications: utils.ConvertNotficationsToProtoFormat(&d.Notifications),
	}
	return res, nil
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

// TODO: GetAll - get only list of fields: ID, Title, Type, ExpiresAt
// TODO: EditNotification - attrs: Notification{}, DocumentID, UserID
// TODO: DeleteNotification - attrs: ID, DocumentID, UserID
