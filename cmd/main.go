package main

import (
	"log"

	"github.com/arjun-saseendran/institute-desk/internal/class"
	"github.com/arjun-saseendran/institute-desk/internal/db"
	"github.com/arjun-saseendran/institute-desk/internal/enrollment"
	"github.com/arjun-saseendran/institute-desk/internal/session"
	"github.com/arjun-saseendran/institute-desk/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	router := gin.Default()

	userService := user.NewUserService(dbConnection)
	userHandler := user.NewUserHandler(userService)
	userHandler.RegisterEndPoints(router)

	sessionService := session.NewSessionService(dbConnection)
	sessionHandler := session.NewSessionHandler(sessionService)
	sessionHandler.RegisterEndPoints(router)

	classService := class.NewClassService(dbConnection)
	classHandler := class.NewClassHandler(classService)
	classHandler.RegisterEndPoints(router)

	enrollmentService := enrollment.NewEnrollmentService(dbConnection)
	enrollmentHandler := enrollment.NewEnrollmetHandler(enrollmentService)
	enrollmentHandler.RegisterEndPoints(router)

	router.Run(":3000")

}
