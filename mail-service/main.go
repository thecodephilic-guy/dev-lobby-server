package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/mail-service/config"
	"github.com/thecodephilic-guy/mail-service/handlers"
)

func main() {
	//connect to DB:
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize handlers with dependencies
	h := handlers.NewHandlerDeps(db)

	router := gin.Default()

	router.POST("/verify/email", h.SendOTPHandler)

	router.Run(":8082")
}
