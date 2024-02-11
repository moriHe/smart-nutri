package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func extractBearerToken(authHeader string) string {
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}
	return ""
}

func (s *Server) GetIdToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return ""
	}

	idToken := extractBearerToken(authHeader)
	if idToken == "" {
		return ""
	}

	user, err := s.auth.Auth.User(c, idToken)
	if err != nil {
		return ""
	}
	return user.ID
}

func (s *Server) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fireUid := s.GetIdToken(c)
		if fireUid == "" {
			responses.ErrorResponse(c, &types.UnauthorizedError)
			return
		}
		user, err := s.store.GetUser(fireUid)
		if err != nil {
			responses.ErrorResponse(c, &types.UnauthorizedError)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
