package api

import (
	"fmt"
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
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	// TODO Check that user is OWNER
	token, err := s.store.GetInvitationLink(user)

	encodedToken := url.QueryEscape(token)
	invitationURL := fmt.Sprintf("%s/einladung/akzeptieren?token=%s", os.Getenv("FRONTEND_URL"), encodedToken)
	responses.HandleResponse(c, invitationURL, err)
}

func (s *Server) handleAcceptInvitation(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	token := c.Query("token")

	responses.HandleResponse(c, "Joined community", s.store.AcceptInvitation(user.Id, token))
}
