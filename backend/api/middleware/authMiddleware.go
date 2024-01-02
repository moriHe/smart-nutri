package middleware

import (
	"net/http"

	"firebase.google.com/go/v4/auth"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
)

func extractBearerToken(authHeader string) string {
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}
	return ""
}

func AuthMiddleware(store storage.Storage, auth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
			c.Abort()
			return
		}

		idToken := extractBearerToken(authHeader)
		if idToken == "" {
			responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
			c.Abort()
			return
		}

		token, err := auth.VerifyIDToken(c, idToken)
		if err != nil {
			responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
			c.Abort()
			return
		}
		fireUid := token.UID
		user, err := store.GetUser(fireUid)

		if err != nil {
			responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
			c.Abort()
			return
		}

		// Set user in the context for further use
		c.Set("user", user)
		c.Next()
	}
}
