package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
	"github.com/nedpals/supabase-go"
)

func (s *Server) userRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("", s.handlePostUser)
	userGroup.Use(s.AuthMiddleWare())
	userGroup.GET("", s.handleGetUser)
	userGroup.GET("/familys", s.handleGetUserFamilys)
	userGroup.PATCH("", s.handlePatchUser)
	userGroup.DELETE("/delete", s.handleDeleteUser)
}

func (s *Server) handleGetUser(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	responses.HandleResponse(c, user, nil)
}

func (s *Server) handlePostUser(c *gin.Context) {
	fireUid := s.GetIdToken(c)
	if fireUid == "" {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	} else {
		userId, err := s.store.PostUser(fireUid)
		responses.HandleResponse(c, userId, err)
	}

}

func (s *Server) handleGetUserFamilys(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	familys, err := s.store.GetUserFamilys(user.Id)
	responses.HandleResponse(c, familys, err)
}

func (s *Server) handlePatchUser(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	var payload types.PatchUser
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.BadRequestError)
		return
	} else {
		err := s.store.PatchUser(user.Id, payload.NewActiveFamilyId)
		responses.HandleResponse(c, "Patch succeeded", err)
	}

}

func injectAuthorizationHeader(req *http.Request, value string) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", value))
}

func (s *Server) handleDeleteUser(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	if user == nil {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return
	}
	supabaseUid := s.GetIdToken(c)

	err := s.store.DeleteUser(user.Id)
	if err != nil {
		responses.ErrorResponse(c, err)
		return
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ADMIN_KEY")
	client := supabase.CreateClient(supabaseUrl, supabaseKey)
	reqURL := fmt.Sprintf("%s/%s/users/%s", client.BaseURL, "auth/v1/admin", supabaseUid)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, reqURL, nil)
	if err != nil {
		return
	}
	injectAuthorizationHeader(req, supabaseKey)
	req.Header.Set("apikey", supabaseKey)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.InternalServerError, "Failed to delete account"))
		return

	}
	responses.HandleResponse(c, "Account deleted", nil)
	defer resp.Body.Close()

}
