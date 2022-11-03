package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/user/internal/models/user"
	proto "github.com/samgozman/validity.red/user/proto"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedAuthServiceServer
}

type UserServer struct {
	App *Config
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedUserServiceServer
}

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAuthServiceServer(s, &AuthServer{
		App: app,
	})
	proto.RegisterUserServiceServer(s, &UserServer{
		App: app,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (us *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*emptypb.Empty, error) {
	input := req.GetRegisterEntry()

	// register user
	userPayload := user.User{
		Email:    input.Email,
		Password: input.Password,
		Timezone: input.Timezone,
	}
	err := us.App.Repo.InsertOne(ctx, &userPayload)
	if err != nil {
		return nil, err
	}

	// return response
	return &emptypb.Empty{}, nil
}

func (as *AuthServer) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	input := req.GetAuthEntry()

	// find user
	u, err := as.App.Repo.FindOne(ctx, &user.User{Email: input.Email}, "id, password, calendar_id, timezone")
	if err != nil {
		return nil, err
	}

	// verify password
	err = user.VerifyPassword(u.Password, input.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, status.Error(codes.Unauthenticated, "authentication failed")
	}

	// return response
	res := &proto.AuthResponse{
		// TODO: Return user entity
		UserId:     u.ID.String(),
		CalendarId: u.CalendarID,
		Timezone:   u.Timezone,
	}
	return res, nil
}

func (us *UserServer) GetCalendarOptions(ctx context.Context, req *proto.GetCalendarIdRequest) (*proto.GetCalendarIdResponse, error) {
	userId, _ := uuid.Parse(req.UserId)
	u, err := us.App.Repo.FindOne(ctx, &user.User{ID: userId}, "calendar_id, iv_calendar, timezone")
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIdResponse{
		CalendarId: u.CalendarID,
		CalendarIv: u.IV_Calendar,
		Timezone:   u.Timezone,
	}
	return res, nil
}

func (us *UserServer) GetCalendarIv(ctx context.Context, req *proto.GetCalendarIvRequest) (*proto.GetCalendarIvResponse, error) {
	u, err := us.App.Repo.FindOne(ctx, &user.User{CalendarID: req.CalendarId}, "iv_calendar")
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIvResponse{
		CalendarIv: u.IV_Calendar,
	}
	return res, nil
}

func (us *UserServer) SetCalendarIv(ctx context.Context, req *proto.SetCalendarIvRequest) (*emptypb.Empty, error) {
	err := us.App.Repo.Update(ctx, req.UserId, map[string]interface{}{
		"iv_calendar": req.CalendarIv,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// TODO: Edit - edit user info
