package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type Config struct {
	logger          *Logger
	usersClient     *UsersClient
	documentsClient *DocumentsClient
}

type UsersClient struct {
	authService user.AuthServiceClient
	userService user.UserServiceClient
}

type DocumentsClient struct {
	documentService     document.DocumentServiceClient
	notificationService document.NotificationServiceClient
}

func main() {
	// Create logger
	logger := Logger{}

	// ! Move client connections to the separate gorutine
	// ! which will be trying to reconnect without blocking the main app

	// USERS CLIENT SECTION - START //
	userServiceConn, err := connectToUserService()
	if err != nil {
		go logger.LogFatal(&logs.Log{
			Service: "user-service",
			Message: "Error on connecting to the user-service",
			Error:   err.Error(),
		})
		return
	}
	defer userServiceConn.Close()

	usersClient := UsersClient{
		authService: user.NewAuthServiceClient(userServiceConn),
		userService: user.NewUserServiceClient(userServiceConn),
	}
	// USERS CLIENT SECTION - END //

	// DOCUMENTS CLIENT SECTION - START //
	documentServiceConn, err := connectToDocumentService()
	if err != nil {
		go logger.LogFatal(&logs.Log{
			Service: "document-service",
			Message: "Error on connecting to the document-service",
			Error:   err.Error(),
		})
		return
	}
	defer documentServiceConn.Close()

	documentsClient := DocumentsClient{
		documentService:     document.NewDocumentServiceClient(documentServiceConn),
		notificationService: document.NewNotificationServiceClient(documentServiceConn),
	}
	// DOCUMENTS CLIENT SECTION - END //

	app := Config{
		logger:          &logger,
		usersClient:     &usersClient,
		documentsClient: &documentsClient,
	}

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("BROKER_PORT")),
		Handler: app.routes(),
	}

	// start http server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
