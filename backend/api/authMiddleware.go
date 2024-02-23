package api

import (
	"github.com/gin-gonic/gin"
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
		supabaseUid := s.GetIdToken(c)
		if supabaseUid == "" {
			return
		}
		user, err := s.store.GetUser(supabaseUid)
		if err != nil {
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
