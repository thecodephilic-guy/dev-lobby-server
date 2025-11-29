package config

import (
	"os"

	"github.com/thecodephilic-guy/dev-lobby-server/models"
)

func LoadSerices() *models.ServiceConfig {
	return &models.ServiceConfig{
		AuthServiceURL: getEndOrDefault("AUTH_SERVICE_URL", "http://localhost:8081"),
		MailServiceURL: getEndOrDefault("MAIL_SERVICE_URL", "http://localhost:8082"),
	}
}

func getEndOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
