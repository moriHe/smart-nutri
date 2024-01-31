package api

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
	"github.com/nedpals/supabase-go"
)

type Server struct {
	store storage.Storage
	auth  *supabase.Client
}

func StartGinServer(store storage.Storage, url string) (*gin.Engine, error) {
	router := gin.Default()
	supabase := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	server := &Server{store: store, auth: supabase}

	config := cors.DefaultConfig()

	// setup before going live
	config.AllowOrigins = []string{"*"} // Update with your Angular app's origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization", "Secret"}

	// Use the CORS middleware
	router.Use(cors.New(config))

	router.GET("/secret", server.handleGetSecret)
	server.userRoutes(router)
	router.Use(server.AuthMiddleWare())
	server.recipeRoutes(router)
	server.mealPlanRoutes(router)
	server.mealplanShoppingListRoutes(router)
	server.familyRoutes(router)
	server.invitationRoutes(router)

	router.Run(url)

	return router, nil

}

func (s *Server) handleGetSecret(c *gin.Context) {
	secret := os.Getenv("USER_SECRET")
	secretHeader := c.GetHeader("Secret")
	if secret == secretHeader {
		responses.HandleResponse[string](c, "Permission acknowledged", nil)
	} else {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Permission denied"})
	}
	c.Abort()
}
