package main

import (
	"context"
	"errors"
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
		name     string
		fields   fields
		args     args
		want     *proto.Response
		wantErr  bool
		errorMsg error
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
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					UserID: "justWrongId",
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
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "justWrongId",
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
		{
			name:   "should fail if documentId is not exists",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "00000000-0000-0000-0000-000000000000",
						Date:       timestamppb.Now(),
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrDocumentNotFound,
		},
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
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("NotificationServer.Create() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
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
		name     string
		fields   fields
		args     args
		want     *proto.Response
		wantErr  bool
		errorMsg error
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
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					UserID: "justWrongId",
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
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "justWrongId",
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
		{
			name:   "should fail if notificationId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
						Date:       timestamppb.Now(),
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidNotificationId,
		},
		{
			name:   "should fail if documentId is not exists",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						ID:         "8e5d4e72-357e-4838-b371-84693f44c4c3",
						DocumentID: "00000000-0000-0000-0000-000000000000",
						Date:       timestamppb.Now(),
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrDocumentNotFound,
		},
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
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("NotificationServer.Edit() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
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
		name     string
		fields   fields
		args     args
		want     *proto.Response
		wantErr  bool
		errorMsg error
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
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					UserID: "justWrongId",
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
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "justWrongId",
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
		{
			name:   "should fail if notificationId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
						Date:       timestamppb.Now(),
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidNotificationId,
		},
		{
			name:   "should fail if documentId is not exists",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationCreateRequest{
					NotificationEntry: &proto.Notification{
						ID:         "8e5d4e72-357e-4838-b371-84693f44c4c3",
						DocumentID: "00000000-0000-0000-0000-000000000000",
						Date:       timestamppb.Now(),
					},
					UserID: "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrDocumentNotFound,
		},
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
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("NotificationServer.Delete() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
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
		name     string
		fields   fields
		args     args
		want     *proto.ResponseNotificationsList
		wantErr  bool
		errorMsg error
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
		{
			name:   "should fail if userId is incorrect",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationsRequest{
					DocumentID: "434377cf-7509-4cc0-9895-0afa683f0e56",
					UserID:     "wrongId",
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
				req: &proto.NotificationsRequest{
					DocumentID: "wrongId",
					UserID:     "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrInvalidDocumentId,
		},
		{
			name:   "should fail if documentId is not exists",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: &proto.NotificationsRequest{
					DocumentID: "00000000-0000-0000-0000-000000000000",
					UserID:     "458c9061-5262-48b7-9b87-e47fa64d654c",
				},
			},
			wantErr:  true,
			errorMsg: ErrDocumentNotFound,
		},
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
			if tt.wantErr && !errors.Is(err, tt.errorMsg) {
				t.Errorf("NotificationServer.GetAll() wrong error msg = %v, want %v", err.Error(), tt.errorMsg.Error())
				return
			}
			if !tt.wantErr && got.Result != tt.want.Result {
				t.Errorf("NotificationServer.GetAll() = %v, want %v", got, tt.want)
			}
			// TODO: Fix reflect.DeepEqual false negative with []*proto.Notification
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NotificationServer.GetAll() = %v, want %v", got, tt.want)
			// }
		})
	}
}
