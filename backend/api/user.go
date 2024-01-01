package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func extractBearerToken(authHeader string) string {
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}
	return ""
}

func (s *Server) userRoutes(r *gin.Engine) {
	r.GET("/user", s.handleGetUser)
	r.POST("/user", s.handlePostUser)
}

func (s *Server) handleGetUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		errorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
		return
	}

	idToken := extractBearerToken(authHeader)
	if idToken == "" {
		errorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
		return
	}

	token, err := s.Auth.VerifyIDToken(c, idToken)
	if err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
		return
	}
	fireUid := token.UID
	user, err := s.store.GetUser(fireUid)
	handleResponse[*types.User](c, user, err)
}

func (s *Server) handlePostUser(c *gin.Context) {
	var payload types.PostUser

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
	} else {
		userId, err := s.store.PostUser(payload)
		handleResponse[*int](c, userId, err)
	}

}
