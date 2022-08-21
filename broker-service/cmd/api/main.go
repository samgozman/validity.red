package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samgozman/validity.red/broker/internal/token"
	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type Config struct {
	logger          *Logger
	token           *token.TokenMaker
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
	loggerServiceConn, err := connectToService("logger-service", os.Getenv("LOGGER_GRPC_PORT"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer loggerServiceConn.Close()
	logger := Logger{
		client: logs.NewLogServiceClient(loggerServiceConn),
	}

	// ! Move client connections to the separate gorutine
	// ! which will be trying to reconnect without blocking the main app

	// USERS CLIENT SECTION - START //
	userServiceConn, err := connectToService("user-service", os.Getenv("USER_GRPC_PORT"))
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
	documentServiceConn, err := connectToService("document-service", os.Getenv("DOCUMENT_GRPC_PORT"))
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

	// Create JWT token maker
	token := token.TokenMaker{
		Key: []byte(os.Getenv("JWT_SECRET")),
		// 10 minutes 10 * 60
		MaxAge: 10 * 60,
	}

	app := Config{
		logger:          &logger,
		token:           &token,
		usersClient:     &usersClient,
		documentsClient: &documentsClient,
	}

	router := app.routes()
	err = router.Run(fmt.Sprintf(":%s", os.Getenv("BROKER_PORT")))
	if err != nil {
		log.Panic(err)
	}
}
