package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/config"
	"github.com/thecodephilic-guy/dev-lobby-server/middlewares"
	"github.com/thecodephilic-guy/dev-lobby-server/proxy"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Validate JWT secret exists
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatal("JWT_SECRET_KEY not set in environment")
	}

	// Load all the services this gateway is connected to
	services := config.LoadSerices()

	// Creating a default router
	router := gin.Default()

	// Health check (public)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "API Gateway running"})
	})

	// PUBLIC ROUTES - No authentication required
	publicRoutes := router.Group("/api")
	{
		publicRoutes.Any("/auth/*proxyPath", proxy.ReverseProxy(services.AuthServiceURL))
	}

	// PROTECTED ROUTES - JWT authentication required
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuth()) // Apply JWT middleware (Bouncer at the gate)
	{
		// Mail service - All routes require authentication
		protectedRoutes.Any("/mail/*proxyPath", proxy.ReverseProxy(services.MailServiceURL))

		// Add more protected services here
	}

	log.Printf("API Gateway starting on :8080")
	log.Printf("Auth Service: %s", services.AuthServiceURL)
	log.Printf("Mail Service: %s", services.MailServiceURL)

	router.Run(":8080")
}
