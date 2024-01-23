package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
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
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	// Use the CORS middleware
	router.Use(cors.New(config))

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
