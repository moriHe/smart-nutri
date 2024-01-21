package api

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) invitationRoutes(r *gin.Engine) {
	r.GET("invitations/link", s.handleGetInvitationLink)
	r.GET("invitations/accept", s.handleAcceptInvitation)
}

func (s *Server) handleGetInvitationLink(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	// TODO Check that user is OWNER
	token, err := s.store.GetInvitationLink(user.ActiveFamilyId)

	if err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Failed to generate link"})
		return
	}
	encodedToken := url.QueryEscape(token)
	invitationURL := fmt.Sprintf("%sinvitations/accept?token=%s", os.Getenv("SERVER_URL"), encodedToken)
	responses.HandleResponse[string](c, invitationURL, nil)
}

func (s *Server) handleAcceptInvitation(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	token := c.Query("token")

	responses.HandleResponse[string](c, "Joined community", s.store.AcceptInvitation(user.Id, token))
}
