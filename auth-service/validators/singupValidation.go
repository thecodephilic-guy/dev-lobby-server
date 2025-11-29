package validators

import (
	"errors"
	"net/mail"

	"github.com/thecodephilic-guy/auth-service/models"
)

func ValidateSignupForm(reqBody models.AuthRequest) error {
	email := reqBody.Email
	password := reqBody.Password

	if len(email) == 0 || len(password) == 0 {
		return errors.New("both the fields are required")
	}

	//validate email:
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}

	//validate password:
	if len(password) < 6 {
		return errors.New("password should be atleast 6 characters long")
	}

	return nil //means data is valid

}
