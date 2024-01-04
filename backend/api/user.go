package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) userRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("", s.handlePostUser)
	userGroup.Use(s.AuthMiddleWare())
	userGroup.GET("", s.handleGetUser)
}

func (s *Server) handleGetUser(c *gin.Context) {
	user, err := contextmethods.GetUserFromContext(c)
	responses.HandleResponse[*types.User](c, user, err)
}

func (s *Server) handlePostUser(c *gin.Context) {
	fireUid := s.GetIdToken(c)
	if fireUid == "" {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized"})
		c.Abort()
		return
	} else {
		userId, err := s.store.PostUser(fireUid)
		responses.HandleResponse[*int](c, userId, err)
	}

}
