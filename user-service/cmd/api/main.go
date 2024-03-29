package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/samgozman/validity.red/user/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Repo user.UserRepository
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 0.2,
		SampleRate:       1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	// Connect to SQL server
	db := connectToDB()
	if db == nil {
		panic("Can't connect to Postgres!")
	}

	//Automatic migration for users table
	err = db.Table("users").AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}

	// Create app
	app := Config{}
	app.setupRepo(db)

	// Start gRPC server
	// go app.gRPCListen()
	app.gRPCListen()
}

func connectToDB() *gorm.DB {
	var counts uint8

	for {
		connection, err := gorm.Open(postgres.New(postgres.Config{
			DSN: getPostgresDSN(),
		}))
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for three seconds...")
		time.Sleep(3 * time.Second)

		continue
	}
}

func (app *Config) setupRepo(conn *gorm.DB) {
	db := user.NewPostgresRepository(conn)
	app.Repo = db
}
