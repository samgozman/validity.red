package main

import (
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/grpc"
)

var gRPCPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterDocumentServiceServer(s, &DocumentServer{
		App: app,
	})
	proto.RegisterNotificationServiceServer(s, &NotificationServer{
		App: app,
	})

	log.Printf("GRPC server listening on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}
