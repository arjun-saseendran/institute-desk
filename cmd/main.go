package main

import (
	"log"

	"github.com/arjun-saseendran/institute/internal/db"
	"github.com/arjun-saseendran/institute/internal/user"
	"github.com/gin-gonic/gin"
)


func main(){
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v",err)
	}
	
	router := gin.Default()
	
	userService := user.NewUserService(dbConnection)
	userHandler := user.NewUserHandler(userService)
	userHandler.RegisterEndPoints(router)
	
	router.Run(":3000")
	
}