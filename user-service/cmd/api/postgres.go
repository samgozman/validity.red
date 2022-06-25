// Helper methods for postgres db

package main

import (
	"fmt"
	"os"
)

func getPostgresDSN() string {
	var (
		dbname = os.Getenv("POSTGRES_DB")
		user   = os.Getenv("POSTGRES_USER")
		pass   = os.Getenv("POSTGRES_PASSWORD")
		host   = os.Getenv("POSTGRES_HOST")
		port   = os.Getenv("POSTGRES_PORT")
	)

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=UTC connect_timeout=5",
		host,
		port,
		user,
		dbname,
		pass,
	)
}
