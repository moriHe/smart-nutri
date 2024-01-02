package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
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

func (s *Server) userRoutes(r *gin.Engine) {
	r.GET("/user", s.handleGetUser)
	r.POST("/user", s.handlePostUser)
}

func (s *Server) handleGetUser(c *gin.Context) {
	user, err := contextmethods.GetUserFromContext(c)
	responses.HandleResponse[*types.User](c, user, err)
}

func (s *Server) handlePostUser(c *gin.Context) {
	var payload types.PostUser

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
	} else {
		userId, err := s.store.PostUser(payload)
		responses.HandleResponse[*int](c, userId, err)
	}

}
