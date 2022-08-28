package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	proto "github.com/samgozman/validity.red/document/proto"
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
		Result: fmt.Sprintf(
			"User '%s' created document '%s' successfully!",
			okReq.DocumentEntry.UserID,
			"00000000-0000-0000-0000-000000000000",
		),
		DocumentId: "00000000-0000-0000-0000-000000000000",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ResponseDocumentCreate
		wantErr bool
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
		// TODO: Add test for wrong arguments
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

	okRes := &proto.Response{
		Result: fmt.Sprintf(
			"User '%s' edited document '%s' successfully!",
			okReq.DocumentEntry.UserID,
			okReq.DocumentEntry.ID,
		),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Response
		wantErr bool
	}{
		{
			name:   "should edit document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		// TODO: Add tests for error cases
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

	okRes := &proto.Response{
		Result: fmt.Sprintf(
			"User '%s' deleted document '%s' successfully!",
			okReq.UserID,
			okReq.DocumentID,
		),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Response
		wantErr bool
	}{
		{
			name:   "should delete document",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		// TODO: Add tests for error cases
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
		Result: fmt.Sprintf(
			"User '%s' found document '%s' successfully!",
			okReq.UserID,
			okReq.DocumentID,
		),
		Document: doc,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ResponseDocument
		wantErr bool
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
		// TODO: Add tests for error cases
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
		Result: fmt.Sprintf(
			"User '%s' found %d documents successfully!",
			okReq.UserID,
			0,
		),
		Documents: []*proto.Document{},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ResponseDocumentsList
		wantErr bool
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
		// TODO: Add tests for error cases
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentServer.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
