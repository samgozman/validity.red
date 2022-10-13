package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/samgozman/validity.red/user/internal/models/user"
	proto "github.com/samgozman/validity.red/user/proto"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc"
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

func (us *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.Response, error) {
	input := req.GetRegisterEntry()

	// register user
	userPayload := user.User{
		Email:    input.Email,
		Password: input.Password,
	}
	err := us.App.Repo.InsertOne(ctx, &userPayload)
	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User '%s' registered successfully!", userPayload.ID)}
	return res, nil
}

func (as *AuthServer) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	input := req.GetAuthEntry()

	// find user
	u, err := as.App.Repo.FindOneByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	// verify password
	err = user.VerifyPassword(u.Password, input.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	// return response
	res := &proto.AuthResponse{
		Result: fmt.Sprintf("User '%s' logged in successfully!", u.ID),
		// TODO: Return user entity
		UserId:     u.ID.String(),
		CalendarId: u.CalendarID,
	}
	return res, nil
}

func (us *UserServer) GetCalendarId(ctx context.Context, req *proto.GetCalendarIdRequest) (*proto.GetCalendarIdResponse, error) {
	u, err := us.App.Repo.GetCalendarId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIdResponse{
		CalendarId: u.CalendarID,
		CalendarIv: u.IV_Calendar,
	}
	return res, nil
}

func (us *UserServer) GetCalendarIv(ctx context.Context, req *proto.GetCalendarIvRequest) (*proto.GetCalendarIvResponse, error) {
	iv, err := us.App.Repo.GetCalendarIv(ctx, req.CalendarId)
	if err != nil {
		return nil, err
	}

	res := &proto.GetCalendarIvResponse{
		CalendarIv: iv,
	}
	return res, nil
}

func (us *UserServer) SetCalendarIv(ctx context.Context, req *proto.SetCalendarIvRequest) (*proto.Response, error) {
	err := us.App.Repo.Update(ctx, req.UserId, map[string]interface{}{
		"iv_calendar": req.CalendarIv,
	})
	if err != nil {
		return nil, err
	}

	res := &proto.Response{
		Result: "Calendars IV updated successfully!",
	}
	return res, nil
}

// TODO: Edit - edit user info
