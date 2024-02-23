package api

import (
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
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	var payload types.FamilyBody

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, err.Error()))
		return
	} else {
		err := s.store.PostFamily(payload.Name, user.Id)
		responses.HandleResponse(c, "Added family", err)
	}
}
