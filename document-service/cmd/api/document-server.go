package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DocumentServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedDocumentServiceServer
}

func (ds *DocumentServer) Create(ctx context.Context, req *proto.DocumentCreateRequest) (*proto.ResponseDocumentCreate, error) {
	input := req.GetDocumentEntry()

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, ErrInvalidUserId
	}

	// register document
	d := document.Document{
		UserID:      userID,
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	}
	err = ds.App.Documents.InsertOne(ctx, &d)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocumentCreate{
		Result:     fmt.Sprintf("User '%s' created document '%s' successfully!", userID, d.ID),
		DocumentId: d.ID.String(),
	}
	return res, nil
}

func (ds *DocumentServer) Edit(ctx context.Context, req *proto.DocumentCreateRequest) (*proto.Response, error) {
	input := req.GetDocumentEntry()

	// Decode values
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, ErrInvalidDocumentId
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, ErrInvalidUserId
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
	err = ds.App.Documents.UpdateOne(ctx, &d)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' edited document '%s' successfully!", userID, d.ID)}
	return res, nil
}

func (ds *DocumentServer) Delete(ctx context.Context, req *proto.DocumentRequest) (*proto.Response, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, ErrInvalidDocumentId
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserId
	}

	// delete document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = ds.App.Documents.DeleteOne(ctx, &d)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' deleted document '%s' successfully!", userID, id)}
	return res, nil
}

func (ds *DocumentServer) GetOne(ctx context.Context, req *proto.DocumentRequest) (*proto.ResponseDocument, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, ErrInvalidDocumentId
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserId
	}

	// Find document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = ds.App.Documents.FindOne(ctx, &d)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocument{
		Result: fmt.Sprintf("User '%s' found document '%s' successfully!", userID, d.ID),
		Document: &proto.Document{
			ID:          d.ID.String(),
			UserID:      d.UserID.String(),
			Title:       d.Title,
			Type:        d.Type,
			Description: d.Description,
			ExpiresAt:   timestamppb.New(d.ExpiresAt),
		},
	}
	return res, nil
}

func (ds *DocumentServer) GetAll(ctx context.Context, req *proto.DocumentsRequest) (*proto.ResponseDocumentsList, error) {
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserId
	}

	// Find all documents
	documents, err := ds.App.Documents.FindAll(ctx, userID)

	// return error if exists
	if err != nil {
		return nil, err
	}

	// Transform documents to proto format
	protoDocuments := make([]*proto.Document, len(documents))
	for i, d := range documents {
		protoDocuments[i] = &proto.Document{
			ID:          d.ID.String(),
			Title:       d.Title,
			Type:        d.Type,
			Description: d.Description,
			ExpiresAt:   timestamppb.New(d.ExpiresAt),
		}
	}

	// return response
	res := &proto.ResponseDocumentsList{
		Result:    fmt.Sprintf("User '%s' found %d documents successfully!", userID, len(documents)),
		Documents: protoDocuments,
	}
	return res, nil
}
