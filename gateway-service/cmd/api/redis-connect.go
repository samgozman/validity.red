package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

// Connect to Redis and return a Redis client.
// Wait for the connection to be established before returning.
func connectToRedis(conf RedisConfig) *redis.Client {
	var counts uint8
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

		if counts > 5 {
			log.Fatalln(err)
			return nil
		}

		time.Sleep(3 * time.Second)
		continue
	}
}
