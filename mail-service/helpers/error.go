package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/mail-service/models"
)

func SendError(ctx *gin.Context, statusCode int, message string, details string) {
	response := models.ErrorResponse{
		StatusCode: statusCode,
		Title:      message,
		Error:      details,
	}

	ctx.IndentedJSON(statusCode, response)
}
