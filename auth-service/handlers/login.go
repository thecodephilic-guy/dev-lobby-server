package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/auth-service/helpers"
	"github.com/thecodephilic-guy/auth-service/models"
	"github.com/thecodephilic-guy/auth-service/validators"
)

func (h *HandlerDeps) Login(ctx *gin.Context) {
	var reqBody models.AuthRequest

	if !helpers.BindJSON(ctx, &reqBody) {
		return //if failed then return
	}

	email := reqBody.Email
	password := reqBody.Password

	if validationError := validators.ValidateLogin(reqBody); validationError != nil {
		helpers.SendError(ctx, http.StatusUnprocessableEntity, "Validation failed", validationError.Error())
		return
	}

	//fetch data using email:
	var dbResponse models.User
	if err := h.DB.QueryRow(ctx, helpers.FindUser, email).Scan(&dbResponse.ID, &dbResponse.Email, &dbResponse.Password, &dbResponse.IsEmailVerified, &dbResponse.CreatedAt); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "Authentication failed", "your email is incorrect")
		return
	}

	//if incorrect password
	if !helpers.IsPasswordCorrect(password, dbResponse.Password) {
		helpers.SendError(ctx, http.StatusUnauthorized, "Authentication failed", "your password is incorrect")
		return
	}

	//generate jwt:
	token, err := helpers.GenerateToken(dbResponse.ID)
	if err != nil {
		fmt.Print("Failed generaring the token \n", err.Error())
	}

	ctx.SetCookie("token", token, 2*36000, "/", "localhost", false, true)

	helpers.SendResponse(ctx, http.StatusOK, "Login successfull", dbResponse)
}
