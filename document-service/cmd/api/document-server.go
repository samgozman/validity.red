package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/utils"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const NumberOfLatestDocuments = 5

type DocumentServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedDocumentServiceServer
}

func (ds *DocumentServer) Create(ctx context.Context, req *proto.DocumentCreateRequest) (*proto.ResponseDocumentCreate, error) {
	input := req.GetDocumentEntry()

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// Check if user has reached the limit of documents
	count, err := ds.App.Documents.Count(ctx, userID)
	if err != nil {
		return nil, err
	}

	if count >= ds.App.limits.MaxDocumentsPerUser {
		return nil, ErrMaxDocumentsLimit
	}

	// register document
	d := document.Document{
		UserID:      userID,
		Title:       input.Title,
		Type:        &input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	}
	err = ds.App.Documents.InsertOne(ctx, &d)

	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocumentCreate{
		DocumentId: d.ID.String(),
	}

	return res, nil
}

func (ds *DocumentServer) Edit(ctx context.Context, req *proto.DocumentCreateRequest) (*emptypb.Empty, error) {
	input := req.GetDocumentEntry()

	// Decode values
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, ErrInvalidDocumentID
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// update document
	d := document.Document{
		ID:          id,
		UserID:      userID,
		Title:       input.Title,
		Type:        &input.Type,
		Description: input.Description,
		ExpiresAt:   input.ExpiresAt.AsTime(),
	}
	err = ds.App.Documents.UpdateOne(ctx, &d)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ds *DocumentServer) Delete(ctx context.Context, req *proto.DocumentRequest) (*emptypb.Empty, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, ErrInvalidDocumentID
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// delete document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = ds.App.Documents.DeleteOne(ctx, &d)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ds *DocumentServer) GetOne(ctx context.Context, req *proto.DocumentRequest) (*proto.ResponseDocument, error) {
	// Decode values
	id, err := uuid.Parse(req.GetDocumentID())
	if err != nil {
		return nil, ErrInvalidDocumentID
	}

	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// Find document
	d := document.Document{
		ID:     id,
		UserID: userID,
	}
	err = ds.App.Documents.FindOne(ctx, &d)

	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocument{
		Document: &proto.Document{
			ID:          d.ID.String(),
			UserID:      d.UserID.String(),
			Title:       d.Title,
			Type:        *d.Type,
			Description: d.Description,
			ExpiresAt:   timestamppb.New(d.ExpiresAt),
		},
	}

	return res, nil
}

func (ds *DocumentServer) GetAll(ctx context.Context, req *proto.DocumentsRequest) (*proto.ResponseDocumentsList, error) {
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserID
	}

	// Find all documents
	documents, err := ds.App.Documents.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.ResponseDocumentsList{
		Documents: utils.ConvertDocumentsToProtoFormat(&documents),
	}

	return res, nil
}

func (ds *DocumentServer) GetUserStatistics(
	ctx context.Context,
	req *proto.DocumentsRequest,
) (*proto.ResponseDocumentsStatistics, error) {
	userID, err := uuid.Parse(req.GetUserID())
	if err != nil {
		return nil, ErrInvalidUserID
	}

	total, err := ds.App.Documents.Count(ctx, userID)
	if err != nil {
		return nil, err
	}

	types, err := ds.App.Documents.CountTypes(ctx, userID)
	if err != nil {
		return nil, err
	}

	latest, err := ds.App.Documents.FindLatest(ctx, userID, NumberOfLatestDocuments)
	if err != nil {
		return nil, err
	}

	return &proto.ResponseDocumentsStatistics{
		Total:           total,
		Types:           types,
		LatestDocuments: utils.ConvertDocumentsToProtoFormat(&latest),
	}, nil
}
