package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

// Connect to Redis and return a Redis client.
// Wait for the connection to be established before returning.
func connectToRedis(conf RedisConfig) *redis.Client {
	const requestTimeout = 3 * time.Second

	const maxRetries = 5

	const retryDelay = 3 * time.Second

	var counts uint8

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	for {
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Password: conf.Password,
			DB:       0,
		})
		_, err := rdb.Ping(ctx).Result()

		if err != nil {
			log.Println("Redis not yet ready...")
			counts++
		} else {
			log.Println("Connected to Redis!")
			return rdb
		}

		if counts > maxRetries {
			log.Fatalln(err)
			return nil
		}

		time.Sleep(retryDelay)

		continue
	}
}
