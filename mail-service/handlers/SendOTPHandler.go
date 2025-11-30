package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/mail-service/helpers"
)

func (h *HandlerDeps) SendOTPHandler(ctx *gin.Context) {
	loggedInUserId := ctx.GetHeader("userId")

	otp, err := helpers.GenerateOTP()
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error generating OTP", err.Error())
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"id":  loggedInUserId,
		"otp": otp,
	})
}
