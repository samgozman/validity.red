package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNotificationServer_Create(t *testing.T) {
	type fields struct {
		App                                    *Config
		UnimplementedNotificationServiceServer proto.UnimplementedNotificationServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.NotificationCreateRequest
	}

	okReq := &proto.NotificationCreateRequest{
		NotificationEntry: &proto.Notification{
			DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
			Date:       timestamppb.Now(),
		},
		UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.Response{
		Result: fmt.Sprintf(
			"User '%s' created notification '%s' successfully!",
			okReq.UserID,
			"00000000-0000-0000-0000-000000000000",
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
			name:   "should create notification",
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
			ds := &NotificationServer{
				App:                                    tt.fields.App,
				UnimplementedNotificationServiceServer: tt.fields.UnimplementedNotificationServiceServer,
			}
			got, err := ds.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotificationServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationServer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationServer_Edit(t *testing.T) {
	type fields struct {
		App                                    *Config
		UnimplementedNotificationServiceServer proto.UnimplementedNotificationServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.NotificationCreateRequest
	}

	okReq := &proto.NotificationCreateRequest{
		NotificationEntry: &proto.Notification{
			ID:         "8e5d4e72-357e-4838-b371-84693f44c4c3",
			DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
			Date:       timestamppb.Now(),
		},
		UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.Response{
		Result: fmt.Sprintf(
			"User '%s' edited notification '%s' successfully!",
			okReq.UserID,
			okReq.NotificationEntry.ID,
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
			name:   "should edited notification",
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
			ds := &NotificationServer{
				App:                                    tt.fields.App,
				UnimplementedNotificationServiceServer: tt.fields.UnimplementedNotificationServiceServer,
			}
			got, err := ds.Edit(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotificationServer.Edit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationServer.Edit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationServer_Delete(t *testing.T) {
	type fields struct {
		App                                    *Config
		UnimplementedNotificationServiceServer proto.UnimplementedNotificationServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.NotificationCreateRequest
	}

	okReq := &proto.NotificationCreateRequest{
		NotificationEntry: &proto.Notification{
			ID:         "8e5d4e72-357e-4838-b371-84693f44c4c3",
			DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
			Date:       timestamppb.Now(),
		},
		UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.Response{
		Result: fmt.Sprintf(
			"User '%s' deleted notification with id '%s' successfully!",
			okReq.UserID,
			okReq.NotificationEntry.ID,
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
			name:   "should delete notification",
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
			ds := &NotificationServer{
				App:                                    tt.fields.App,
				UnimplementedNotificationServiceServer: tt.fields.UnimplementedNotificationServiceServer,
			}
			got, err := ds.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotificationServer.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotificationServer.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationServer_GetAll(t *testing.T) {
	type fields struct {
		App                                    *Config
		UnimplementedNotificationServiceServer proto.UnimplementedNotificationServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.NotificationsRequest
	}

	okReq := &proto.NotificationsRequest{
		DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
		UserID:     "458c9061-5262-48b7-9b87-e47fa64d654c",
	}

	okRes := &proto.ResponseNotificationsList{
		Result: fmt.Sprintf(
			"User '%s' found %d notifications successfully!",
			okReq.UserID,
			0,
		),
		// TODO: Fix reflect.DeepEqual false negative
		// Notifications: []*proto.Notification{},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ResponseNotificationsList
		wantErr bool
	}{
		{
			name:   "should find all notifications",
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
			ds := &NotificationServer{
				App:                                    tt.fields.App,
				UnimplementedNotificationServiceServer: tt.fields.UnimplementedNotificationServiceServer,
			}
			got, err := ds.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotificationServer.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Result != tt.want.Result {
				t.Errorf("NotificationServer.GetAll() = %v, want %v", got, tt.want)
			}
			// TODO: Fix reflect.DeepEqual false negative with []*proto.Notification
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NotificationServer.GetAll() = %v, want %v", got, tt.want)
			// }
		})
	}
}
