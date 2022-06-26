package main

import (
	"github.com/samgozman/validity.red/document/internal/models/document"
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

	//Automatic migration for documents table
	err = db.Table("documents").AutoMigrate(&document.Document{})
	if err != nil {
		panic(err)
	}

	// Create app
	app := Config{
		db: db,
	}

	// Start gRPC server
	// app.gRPCListen()
}
