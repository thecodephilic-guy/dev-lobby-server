package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/mail-service/helpers"
	"github.com/thecodephilic-guy/mail-service/models"
)

func (h *HandlerDeps) SendOTPHandler(ctx *gin.Context) {
	//1. extract loggedIn userID
	loggedInUserId := ctx.GetHeader("userId")

	//2. Check if user exist in DB also extract his email:
	var dbResponse models.User
	if err := h.DB.QueryRow(ctx, helpers.SelectUserById, loggedInUserId).Scan(&dbResponse.ID, &dbResponse.Email, &dbResponse.Password, &dbResponse.IsEmailVerified, &dbResponse.CreatedAt); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "User not found", err.Error())
		return
	}

	//3. Check if user is already verified:
	if dbResponse.IsEmailVerified {
		helpers.SendError(ctx, http.StatusBadRequest, "Not a valid request", "User is already verified")
		return
	}

	//4. Generate the OTP
	otp, err := helpers.GenerateOTP()
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Error generating OTP", err.Error())
		return
	}

	//5 Store the OTP into DB:
	expiresAt := time.Now().Add(5 * time.Minute) //otp expires in 5 min
	var response models.Verification
	if err := h.DB.QueryRow(ctx, helpers.InsertNewOTP, loggedInUserId, otp, expiresAt).Scan(&response.ID, &response.UserID, &response.Attempts, &response.ExpiresAt, &response.CreatedAt); err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "Failed to store OTP", err.Error())
		return
	}

	//6. Draft and send mail:
	helpers.SendEmail(dbResponse.Email, otp)

	//7. Send success response
	helpers.SendResponse(ctx, http.StatusOK, "OTP sent successfully", response)
}
