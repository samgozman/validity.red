package main

import (
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/samgozman/validity.red/document/proto"

	"google.golang.org/grpc"
)

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterDocumentServiceServer(s, &DocumentServer{
		db: app.db,
	})
	proto.RegisterNotificationServiceServer(s, &NotificationServer{
		db: app.db,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}
