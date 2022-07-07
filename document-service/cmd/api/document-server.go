package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/utils"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type DocumentServer struct {
	db *gorm.DB
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedDocumentServiceServer
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
