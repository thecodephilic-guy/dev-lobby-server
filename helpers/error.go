package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/models"
)

func SendError(ctx *gin.Context, statusCode int, title string, details string) {
	respones := models.ErrorResponse{
		StatusCode: statusCode,
		Title:      title,
		Error:      details,
	}

	ctx.IndentedJSON(statusCode, respones)
}
