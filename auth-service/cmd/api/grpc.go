package main

import (
	auth "auth/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type AuthServer struct {
	// Necessary parameter to insure backwards compatibility
	auth.UnimplementedAuthServiceServer
}

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	auth.RegisterAuthServiceServer(s, &AuthServer{})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (l *AuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	input := req.GetRegisterEntry()

	// register user
	// return error if exists

	// return response
	res := &auth.RegisterResponse{Result: fmt.Sprintf("User with email %s registered successfully!", input.Email)}
	return res, nil
}
