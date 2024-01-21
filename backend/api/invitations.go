package api

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
)

func (s *Server) invitationRoutes(r *gin.Engine) {
	r.GET("invitations/link", s.handleGetInvitationLink)
	r.GET("invitations/accept", s.handleAcceptInvitation)
}

func (s *Server) handleGetInvitationLink(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	// TODO Check that user is OWNER
	token, err := s.store.GetInvitationLink(user)

	encodedToken := url.QueryEscape(token)
	invitationURL := fmt.Sprintf("%sinvitations/accept?token=%s", os.Getenv("SERVER_URL"), encodedToken)
	responses.HandleResponse[string](c, invitationURL, err)
}

func (s *Server) handleAcceptInvitation(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	token := c.Query("token")

	responses.HandleResponse[string](c, "Joined community", s.store.AcceptInvitation(user.Id, token))
}
