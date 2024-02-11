package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) familyRoutes(r *gin.Engine) {
	r.POST("/familys", s.handlePostFamily)
}

func (s *Server) handlePostFamily(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	var payload types.FamilyBody

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
		return
	} else {
		err := s.store.PostFamily(payload.Name, user.Id)
		responses.HandleResponse(c, "Added family", err)
	}
}
