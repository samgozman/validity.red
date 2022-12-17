package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const BackOffDuration = 2 * time.Second
const BackOffAttempts = 10

func connectToService(name, port string) (*grpc.ClientConn, error) {
	url := fmt.Sprintf("%s:%s", name, port)

	var counts uint8

	for {
		ctxDial, cancelDial := context.WithTimeout(context.Background(), time.Second)
		defer cancelDial()

		conn, err := grpc.DialContext(ctxDial, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

		if err != nil {
			log.Printf("Service '%s' not yet ready...\n", name)
			counts++
		} else {
			log.Printf("Connected to '%s'!\n", name)
			return conn, nil
		}

		if counts > BackOffAttempts {
			return nil, err
		}

		log.Printf("Backing off connection to '%s' for two seconds...", name)
		time.Sleep(BackOffDuration)

		continue
	}
}
