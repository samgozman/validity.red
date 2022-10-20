package main

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/samgozman/validity.red/user/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUserServer_Register(t *testing.T) {
	type fields struct {
		App                            *Config
		UnimplementedUserServiceServer proto.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.RegisterRequest
	}

	okReq := &proto.RegisterRequest{
		RegisterEntry: &proto.Register{
			Email:    "me@example.com",
			Password: "",
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{
			name:   "should register without error",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserServer{
				App:                            tt.fields.App,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := us.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServer.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServer.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthServer_Login(t *testing.T) {
	type fields struct {
		App                            *Config
		UnimplementedAuthServiceServer proto.UnimplementedAuthServiceServer
	}
	type args struct {
		ctx context.Context
		req *proto.AuthRequest
	}

	okReq := &proto.AuthRequest{
		AuthEntry: &proto.Auth{
			Email:    "me@example.com",
			Password: "",
		},
	}

	okRes := &proto.AuthResponse{
		UserId: "434377cf-7509-4cc0-9895-0afa683f0e56",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.AuthResponse
		wantErr bool
	}{
		{
			name:   "should login without error",
			fields: fields{App: &testApp},
			args: args{
				ctx: context.Background(),
				req: okReq,
			},
			want:    okRes,
			wantErr: false,
		},
		// TODO: Test for correct error in case of wrong email or password
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &AuthServer{
				App:                            tt.fields.App,
				UnimplementedAuthServiceServer: tt.fields.UnimplementedAuthServiceServer,
			}
			got, err := us.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
