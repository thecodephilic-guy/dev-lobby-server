package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/models"
)

func SendResponse(ctx *gin.Context, statusCode int, title string, data any) {
	response := models.SuccessResponse{
		StatusCode: statusCode,
		Title:      title,
		Data:       data,
	}

	ctx.IndentedJSON(statusCode, response)
}
