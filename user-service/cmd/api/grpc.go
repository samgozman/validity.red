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
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type AuthServer struct {
	db *gorm.DB
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

	proto.RegisterAuthServiceServer(s, &AuthServer{
		db: app.db,
	})
	proto.RegisterUserServiceServer(s, &UserServer{
		db: app.db,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (us *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.Response, error) {
	input := req.GetRegisterEntry()

	// register user
	err := user.InsertOne(ctx, us.db, &user.User{
		Password: input.Password,
		Email:    input.Email,
	})
	// return error if exists
	if err != nil {
		return nil, err
	}

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User with email %s registered successfully!", input.Email)}
	return res, nil
}

func (us *AuthServer) Login(ctx context.Context, req *proto.AuthRequest) (*proto.Response, error) {
	input := req.GetAuthEntry()

	// find user
	u, err := user.FindOneByEmail(ctx, us.db, &user.User{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}

	// verify password
	err = user.VerifyPassword(u.Password, input.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	// TODO: create JWT token
	// TODO: return JWT token

	// return response
	res := &proto.Response{Result: fmt.Sprintf("User with email %s logged in successfully!", input.Email)}
	return res, nil
}
