package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samgozman/validity.red/broker/internal/token"
	"github.com/samgozman/validity.red/broker/proto/calendar"
	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type Config struct {
	token           *token.TokenMaker
	usersClient     *UsersClient
	documentsClient *DocumentsClient
	calendarsClient *CalendarsClient
}

type UsersClient struct {
	authService user.AuthServiceClient
	userService user.UserServiceClient
}

type DocumentsClient struct {
	documentService     document.DocumentServiceClient
	notificationService document.NotificationServiceClient
}

type CalendarsClient struct {
	calendarService calendar.CalendarServiceClient
}

func main() {
	// ! Move client connections to the separate gorutine
	// ! which will be trying to reconnect without blocking the main app

	// USERS CLIENT SECTION - START //
	userServiceConn, err := connectToService("user-service", os.Getenv("USER_GRPC_PORT"))
	if err != nil {
		log.Fatalln("Error on connecting to the user-service:", err)
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
		log.Fatalln("Error on connecting to the document-service:", err)
		return
	}
	defer documentServiceConn.Close()

	documentsClient := DocumentsClient{
		documentService:     document.NewDocumentServiceClient(documentServiceConn),
		notificationService: document.NewNotificationServiceClient(documentServiceConn),
	}
	// DOCUMENTS CLIENT SECTION - END //

	// CALENDARS CLIENT SECTION - START //
	calendarServiceConn, err := connectToService("calendar-service", os.Getenv("CALENDAR_GRPC_PORT"))
	if err != nil {
		log.Fatalln("Error on connecting to the calendar-service:", err)
		return
	}
	defer calendarServiceConn.Close()

	calendarsClient := CalendarsClient{
		calendarService: calendar.NewCalendarServiceClient(calendarServiceConn),
	}
	// CALENDARS CLIENT SECTION - END //

	// Create JWT token maker
	token := token.TokenMaker{
		Key: []byte(os.Getenv("JWT_SECRET")),
		// 10 minutes 10 * 60
		MaxAge: 10 * 60,
	}

	app := Config{
		token:           &token,
		usersClient:     &usersClient,
		documentsClient: &documentsClient,
		calendarsClient: &calendarsClient,
	}

	router := app.routes()
	err = router.Run(fmt.Sprintf(":%s", os.Getenv("BROKER_PORT")))
	if err != nil {
		log.Panic(err)
	}
}
