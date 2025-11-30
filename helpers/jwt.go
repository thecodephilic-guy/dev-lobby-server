package helpers

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyAndParseToken(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil //token in valid
}

func AppendClaims(ctx *gin.Context, token *jwt.Token) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//append into headers of the req:
		if id, exists := claims["userId"]; exists {
			ctx.Request.Header.Set("userId", id.(string))
		}

		ctx.Set("user_claims", claims)
	}
}
