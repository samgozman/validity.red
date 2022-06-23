package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	user "user/proto"

	"google.golang.org/grpc"
)

type AuthServer struct {
	// Necessary parameter to insure backwards compatibility
	user.UnimplementedAuthServiceServer
}

type UserServer struct {
	// Necessary parameter to insure backwards compatibility
	user.UnimplementedUserServiceServer
}

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	user.RegisterAuthServiceServer(s, &AuthServer{})
	user.RegisterUserServiceServer(s, &UserServer{})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (l *UserServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.Response, error) {
	input := req.GetRegisterEntry()

	// register user
	// return error if exists

	// return response
	res := &user.Response{Result: fmt.Sprintf("User with email %s registered successfully!", input.Email)}
	return res, nil
}
