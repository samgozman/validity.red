package main

import (
	"context"
	"fmt"
	"time"

	"github.com/samgozman/validity.red/broker/proto/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Logger struct {
	client logs.LogServiceClient
}

// TODO: Refactor this to reduce duplications

func (l *Logger) LogDebug(log *logs.Log) (*logs.LogResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := l.client.LogDebug(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *Logger) LogFatal(log *logs.Log) (*logs.LogResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := l.client.LogFatal(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *Logger) LogInfo(log *logs.Log) (*logs.LogResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := l.client.LogInfo(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *Logger) LogWarn(log *logs.Log) (*logs.LogResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := l.client.LogWarn(ctx, &logs.LogRequest{
		LogEntry: log,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
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
