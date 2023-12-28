package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) userRoutes(r *gin.Engine) {
	r.POST("/user", s.handlePostUser)
}

// var payload types.PostRecipe

// 	if err := c.BindJSON(&payload); err != nil {
// 		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
// 	} else {
// 		response, err := s.store.PostRecipe(familyId, payload)
// 		handleResponse[*types.Id](c, response, err)
// 	}

func (s *Server) handlePostUser(c *gin.Context) {
	var payload types.PostUser

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
	} else {
		err := s.store.PostUser(payload)
		handleResponse[string](c, "Hello World", err)
	}

}
