package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/config"
	"github.com/thecodephilic-guy/dev-lobby-server/proxy"
)

func main() {
	//Load all the services this gateway is connected to:
	services := config.LoadSerices()
	//creating a default router:
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "API Gateway running"})
	})

	router.Any("/api/auth/*proxyPath", proxy.ReverseProxy(services.AuthServiceURL))
	// router.Any("/api/mail/*proxyPath", proxy.ReverseProxy(services.MailServiceURL))

	router.Run(":8080")
}
