package validators

import (
	"errors"
	"net/mail"

	"github.com/thecodephilic-guy/auth-service/models"
)

func ValidateLogin(reqBody models.AuthRequest) error {
	email := reqBody.Email
	password := reqBody.Password

	if len(email) == 0 || len(password) == 0 {
		return errors.New("both fields are required")
	}

	//validate email:
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}
