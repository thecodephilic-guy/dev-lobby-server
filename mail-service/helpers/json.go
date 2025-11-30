package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/mail-service/models"
)

// BindJSON is a helper function that binds JSON request body to a struct
// Returns true if binding was successful, false otherwise
// Automatically sends error response if binding fails
func BindJSON(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		errResponse := models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Title:      "Invalid request body",
			Error:      err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, errResponse)
		return false
	}
	return true
}
