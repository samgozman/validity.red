package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	logModel "github.com/samgozman/validity.red/logger/internal/models/log"
	proto "github.com/samgozman/validity.red/logger/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type LogServer struct {
	db *mongo.Database
	// Necessary parameter to insure backwards compatibility
	proto.UnimplementedLogServiceServer
}

var gRpcPort = os.Getenv("GRPC_PORT")

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterLogServiceServer(s, &LogServer{
		db: app.db,
	})

	log.Printf("GRPC server listening on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC: %v", err)
	}
}

func (l *LogServer) LogDebug(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	return logData(ctx, l, req, "DEBUG")
}

func (l *LogServer) LogInfo(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	return logData(ctx, l, req, "INFO")
}

func (l *LogServer) LogWarn(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	return logData(ctx, l, req, "WARN")
}

func (l *LogServer) LogFatal(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	return logData(ctx, l, req, "FATAL")
}

// Log everything based on logLevel
func logData(ctx context.Context, l *LogServer, req *proto.LogRequest, logLevel string) (*proto.LogResponse, error) {
	input := req.GetLogEntry()

	err := logModel.InsertOne(ctx, l.db, logModel.Log{
		Service:  input.Service,
		LogLevel: logLevel,
		Message:  input.Message,
		Error:    input.Error,
	})
	if err != nil {
		return nil, err
	}

	res := &proto.LogResponse{Result: "success"}
	return res, nil
}
