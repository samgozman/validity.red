package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/getsentry/sentry-go"
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

var gRPCPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAuthServiceServer(s, &AuthServer{
		App: app,
	})
	proto.RegisterUserServiceServer(s, &UserServer{
		App: app,
	})

	log.Printf("GRPC server listening on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		sentry.CaptureException(err)
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

// Register - creates a new user and returns it's entity.
func (us *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
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
	return &proto.RegisterResponse{
		UserId: userPayload.ID.String(),
	}, nil
}

// Login - verifies user credentials and returns it's entity.
func (as *AuthServer) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	input := req.GetAuthEntry()

	// find user
	u, err := as.App.Repo.FindOne(ctx, &user.User{Email: input.Email}, "id, password, calendar_id, timezone, is_verified")
	if err != nil {
		return nil, err
	}

	// verify password
	err = user.VerifyPassword(u.Password, input.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, status.Error(codes.Unauthenticated, "authentication failed")
	}

	// return response
	res := &proto.AuthResponse{
		// TODO: Return user entity
		UserId:     u.ID.String(),
		CalendarId: u.CalendarID,
		Timezone:   u.Timezone,
		IsVerified: u.IsVerified,
	}

	return res, nil
}

// GetCalendarOptions gets the calendar field for the user with the given id.
func (us *UserServer) GetCalendarOptions(ctx context.Context, req *proto.GetCalendarIdRequest) (*proto.GetCalendarIdResponse, error) {
	userID, _ := uuid.Parse(req.UserId)

	u, err := us.App.Repo.FindOne(ctx, &user.User{ID: userID}, "calendar_id, iv_calendar, timezone")
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIdResponse{
		CalendarId: u.CalendarID,
		CalendarIv: u.IVCalendar,
		Timezone:   u.Timezone,
	}

	return res, nil
}

// GetCalendarIv gets the iv_calendar field for the user with the given id.
func (us *UserServer) GetCalendarIv(ctx context.Context, req *proto.GetCalendarIvRequest) (*proto.GetCalendarIvResponse, error) {
	u, err := us.App.Repo.FindOne(ctx, &user.User{CalendarID: req.CalendarId}, "iv_calendar")
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIvResponse{
		CalendarIv: u.IVCalendar,
	}

	return res, nil
}

// SetCalendarIv sets the iv_calendar field for the user with the given id.
func (us *UserServer) SetCalendarIv(ctx context.Context, req *proto.SetCalendarIvRequest) (*emptypb.Empty, error) {
	err := us.App.Repo.Update(ctx, req.UserId, map[string]interface{}{
		"iv_calendar": req.CalendarIv,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SetIsVerified sets the is_verified field to true for the user with the given id.
func (us *UserServer) SetIsVerified(ctx context.Context, req *proto.SetIsVerifiedRequest) (*emptypb.Empty, error) {
	err := us.App.Repo.Update(ctx, req.UserId, map[string]interface{}{
		"is_verified": req.IsVerified,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// TODO: Combine all set & get methods
// TODO: Edit - edit user info
