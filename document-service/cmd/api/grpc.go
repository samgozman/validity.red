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
	proto "github.com/samgozman/validity.red/document/proto"
	"gorm.io/gorm"

	"google.golang.org/grpc"
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

	userID, err := uuid.FromBytes(input.UserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// register document
	err = document.InsertOne(ctx, ds.db, &document.Document{
		UserID:      userID,
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	})

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("Document with title '%s' created successfully!", input.Title)}
	return res, nil
}
