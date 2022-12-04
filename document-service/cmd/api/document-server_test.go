package main

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDocumentServer_Create(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentCreateRequest
	}

	okReq := &proto.DocumentCreateRequest{
		DocumentEntry: &proto.Document{
			UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
		},
	}

	okRes := &proto.ResponseDocumentCreate{
		DocumentId: "00000000-0000-0000-0000-000000000000",
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *proto.ResponseDocumentCreate
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should create document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentCreateRequest{
					DocumentEntry: &proto.Document{
						UserID: "justWrongId",
					},
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.Create() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentServer_Edit(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentCreateRequest
	}

	okReq := &proto.DocumentCreateRequest{
		DocumentEntry: &proto.Document{
			ID:     "434377cf-7509-4cc0-9895-0afa683f0e56",
			UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
			Title:  "Edit title",
		},
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *emptypb.Empty
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should edit document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentCreateRequest{
					DocumentEntry: &proto.Document{
						UserID: "justWrongId",
						ID:     "465fef9a-da94-4106-a7b2-83f1f0c2240c",
					},
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
		{
			name:   "should fail if documentId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentCreateRequest{
					DocumentEntry: &proto.Document{
						UserID: "7df12006-b31c-441b-9063-a01ba884b77d",
						ID:     "justWrongId",
					},
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.Edit(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.Edit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.Edit() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.Edit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentServer_Delete(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentRequest
	}

	okReq := &proto.DocumentRequest{
		DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
		UserID:     "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *emptypb.Empty
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should delete document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentRequest{
					DocumentID: "a09f7658-19f7-4854-ab1d-e1052fa2ecce",
					UserID:     "justWrongId",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
		{
			name:   "should fail if documentId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentRequest{
					DocumentID: "justWrongId",
					UserID:     "76193e99-7451-408a-91b8-215761414775",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.Delete() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentServer_GetOne(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentRequest
	}

	doc := &proto.Document{
		ID:        "434377cf-7509-4cc0-9895-0afa683f0e56",
		UserID:    "458c9061-5262-48b7-9b87-e47fa64d654c",
		ExpiresAt: timestamppb.New(time.Time{}),
	}

	okReq := &proto.DocumentRequest{
		DocumentID: doc.ID,
		UserID:     doc.UserID,
	}

	okRes := &proto.ResponseDocument{
		Document: doc,
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *proto.ResponseDocument
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should find document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentRequest{
					DocumentID: "a09f7658-19f7-4854-ab1d-e1052fa2ecce",
					UserID:     "justWrongId",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
		{
			name:   "should fail if documentId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentRequest{
					DocumentID: "justWrongId",
					UserID:     "76193e99-7451-408a-91b8-215761414775",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.GetOne(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.GetOne() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentServer_GetAll(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentsRequest
	}

	okReq := &proto.DocumentsRequest{
		UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.ResponseDocumentsList{
		Documents: []*proto.Document{},
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *proto.ResponseDocumentsList
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should find all documents",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentsRequest{
					UserID: "justWrongId",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.GetAll() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentServer_GetUserStatistics(t *testing.T) {
	type fields struct {
		App                                *Config
		UnimplementedDocumentServiceServer proto.UnimplementedDocumentServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.DocumentsRequest
	}

	okReq := &proto.DocumentsRequest{
		UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.ResponseDocumentsStatistics{
		LatestDocuments: []*proto.Document{},
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *proto.ResponseDocumentsStatistics
		wantErr  bool
		errorMsg error
	}{
		{
			name:   "should find all documents",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.DocumentsRequest{
					UserID: "justWrongId",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidUserId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DocumentServer{
				App:                                tt.fields.App,
				UnimplementedDocumentServiceServer: tt.fields.UnimplementedDocumentServiceServer,
			}
			got, err := ds.GetUserStatistics(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentServer.GetUserStatistics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("DocumentServer.GetUserStatistics() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.GetUserStatistics() = %v, want %v", got, tt.want)
			}
		})
	}
}
