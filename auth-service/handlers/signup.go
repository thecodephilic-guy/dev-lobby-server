package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/auth-service/helpers"
	"github.com/thecodephilic-guy/auth-service/models"
	"github.com/thecodephilic-guy/auth-service/validators"
)

func (h *HandlerDeps) Signup(ctx *gin.Context) {
	var reqBody models.AuthRequest

	if !helpers.BindJSON(ctx, &reqBody) {
		return
	}

	validationError := validators.ValidateSignupForm(reqBody)
	if validationError != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Validation Failed", validationError.Error())
		return
	}

	//Hash the password:
	hashedPassword := helpers.HashPassword(reqBody.Password)

	var response models.User

	err := h.DB.QueryRow(context.Background(), helpers.InsertNewUser, reqBody.Email, hashedPassword).Scan(&response.ID, &response.Email, &response.IsEmailVerified, &response.CreatedAt)
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Operation failed", err.Error())
		return
	}

	//generate jwt:
	token, err := helpers.GenerateToken(response.ID)
	if err != nil {
		fmt.Print("Failed generaring the token \n", err.Error())
	}

	ctx.SetCookie("token", token, 2*3600, "/", "localhost", false, true)

	helpers.SendResponse(ctx, http.StatusCreated, "New user registered", response)
}
