package main

import (
	"github.com/samgozman/validity.red/user/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}

func main() {
	// Connect to SQL server
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: getPostgresDSN(),
	}))
	if err != nil {
		panic(err)
	}

	//Automatic migration for users table
	err = db.Table("users").AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}

	// Create app
	app := Config{
		db: db,
	}

	// Start gRPC server
	// go app.gRPCListen()
	app.gRPCListen()
}
