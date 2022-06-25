package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/samgozman/validity.red/user/internal/models/user"
	proto "github.com/samgozman/validity.red/user/proto"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type AuthServer struct {
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedAuthServiceServer
}

type UserServer struct {
	db *gorm.DB
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

	proto.RegisterAuthServiceServer(s, &AuthServer{})
	proto.RegisterUserServiceServer(s, &UserServer{
		db: app.db,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (u *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.Response, error) {
	input := req.GetRegisterEntry()

	// register user
	err := user.InsertOne(ctx, u.db, &user.User{
		Email: input.Email,
	})
	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User with email %s registered successfully!", input.Email)}
	return res, nil
}
