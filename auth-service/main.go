package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/thecodephilic-guy/auth-service/config"
	"github.com/thecodephilic-guy/auth-service/handlers"
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

	router.POST("/signup", h.Signup)
	router.GET("/login", h.Login)

	router.Run(":8081")
}
