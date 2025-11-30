package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thecodephilic-guy/dev-lobby-server/helpers"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			helpers.SendError(ctx, http.StatusUnauthorized, "Cookie not found", err.Error())
			ctx.Abort()
			return
		}

		//Check format : expected format: header.payload.signature
		partsOfToken := strings.Split(tokenString, ".")
		if len(partsOfToken) != 3 {
			helpers.SendError(ctx, http.StatusBadRequest, "Invalid jwt format", "The jwt token must consist of header, payload, signature")
			ctx.Abort()
			return
		}

		//verify the token:
		token, err := helpers.VerifyAndParseToken(tokenString)
		if err != nil {
			helpers.SendError(ctx, http.StatusUnauthorized, "Unauthorized request", err.Error())
			ctx.Abort()
			return
		}

		//Now if token is valid then extract info it is carrying and add into header
		helpers.AppendClaims(ctx, token)

		//call next handler
		ctx.Next()
	}
}
