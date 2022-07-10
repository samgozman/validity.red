package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/samgozman/validity.red/broker/proto/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Logger struct{}

// TODO: Refactor this to reduce duplications

func (*Logger) LogDebug(log *logs.Log) (*logs.LogResponse, error) {
	ctx, client, conn, cancel, err := connectToLogger()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// call service
	res, err := client.LogDebug(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (*Logger) LogFatal(log *logs.Log) (*logs.LogResponse, error) {
	ctx, client, conn, cancel, err := connectToLogger()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// call service
	res, err := client.LogFatal(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (*Logger) LogInfo(log *logs.Log) (*logs.LogResponse, error) {
	ctx, client, conn, cancel, err := connectToLogger()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// call service
	res, err := client.LogInfo(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (*Logger) LogWarn(log *logs.Log) (*logs.LogResponse, error) {
	ctx, client, conn, cancel, err := connectToLogger()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// call service
	res, err := client.LogWarn(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Function that helps to connect to the loggger service via gRPC protocol
func connectToLogger() (context.Context, logs.LogServiceClient, *grpc.ClientConn, context.CancelFunc, error) {
	// connect to gRPC
	authURL := fmt.Sprintf("logger-service:%s", os.Getenv("LOGGER_GRPC_PORT"))
	conn, err := grpc.Dial(authURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		// app.errorJSON(w, err)
		return nil, nil, nil, nil, err
	}

	// create client
	client := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return ctx, client, conn, cancel, nil
}

func connectToService(name, port string) (*grpc.ClientConn, error) {
	url := fmt.Sprintf("%s:%s", name, port)
	ctxDial, cancelDial := context.WithTimeout(context.Background(), time.Second)
	defer cancelDial()

	conn, err := grpc.DialContext(ctxDial, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
