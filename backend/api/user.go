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
	userGroup.GET("/familys", s.handleGetUserFamilys)
	userGroup.PATCH("", s.handlePatchUser)
}

func (s *Server) handleGetUser(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	responses.HandleResponse(c, user, nil)
}

func (s *Server) handlePostUser(c *gin.Context) {
	fireUid := s.GetIdToken(c)
	if fireUid == "" {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized 1"})
		return
	} else {
		userId, err := s.store.PostUser(fireUid)
		responses.HandleResponse(c, userId, err)
	}

}

func (s *Server) handleGetUserFamilys(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	familys, err := s.store.GetUserFamilys(user.Id)
	responses.HandleResponse(c, familys, err)
}

func (s *Server) handlePatchUser(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	var payload types.PatchUser
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
		return
	} else {
		err := s.store.PatchUser(user.Id, payload.NewActiveFamilyId)
		responses.HandleResponse(c, "Patch succeeded", err)
	}

}
