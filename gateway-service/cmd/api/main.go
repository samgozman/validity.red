package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
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
	redisClient     *redis.Client
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
	// USERS CLIENT SECTION - START //
	userServiceConn, err := connectToService(os.Getenv("USER_GRPC_HOST"), os.Getenv("USER_GRPC_PORT"))
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
	documentServiceConn, err := connectToService(os.Getenv("DOCUMENT_GRPC_HOST"), os.Getenv("DOCUMENT_GRPC_PORT"))
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
	calendarServiceConn, err := connectToService(os.Getenv("CALENDAR_GRPC_HOST"), os.Getenv("CALENDAR_GRPC_PORT"))
	if err != nil {
		log.Fatalln("Error on connecting to the calendar-service:", err)
		return
	}
	defer calendarServiceConn.Close()

	calendarsClient := CalendarsClient{
		calendarService: calendar.NewCalendarServiceClient(calendarServiceConn),
	}
	// CALENDARS CLIENT SECTION - END //

	// Redis connection
	rdb := connectToRedis(RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

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
		redisClient:     rdb,
	}

	router := app.routes()
	err = router.Run(fmt.Sprintf(":%s", os.Getenv("GATEWAY_PORT")))
	if err != nil {
		log.Panic(err)
	}
}
