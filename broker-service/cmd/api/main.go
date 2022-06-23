package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct{}

func main() {
	app := Config{}

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("BROKER_PORT")),
		Handler: app.routes(),
	}

	// start http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
