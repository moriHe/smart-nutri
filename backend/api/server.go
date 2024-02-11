package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

func getOrigin() string {
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		return "*"
	}
	return os.Getenv("FRONTEND_URL")
}

func StartGinServer(store storage.Storage, port string) (*gin.Engine, error) {
	router := gin.Default()
	supabase := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	server := &Server{store: store, auth: supabase}

	// Define the /health endpoint with the specific CORS configuration

	config := cors.DefaultConfig()
	// setup before going live
	config.AllowOrigins = []string{getOrigin()} // Update with your Angular app's origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization", "Secret"}

	// Use the CORS middleware
	router.Use(cors.New(config))
	router.GET("/", server.handleHealthCheck)
	router.GET("/secret", server.handleGetSecret)
	router.GET("/datenbank-nahrungsmittel", server.handleGetIngredientTable)
	server.userRoutes(router)
	router.Use(server.AuthMiddleWare())
	server.recipeRoutes(router)
	server.mealPlanRoutes(router)
	server.mealplanShoppingListRoutes(router)
	server.familyRoutes(router)
	server.invitationRoutes(router)

	router.Run(":" + port)

	return router, nil

}

func (s *Server) handleGetSecret(c *gin.Context) {
	secret := os.Getenv("USER_SECRET")
	secretHeader := c.GetHeader("Secret")
	if secret == secretHeader {
		responses.HandleResponse(c, "Permission acknowledged", nil)
	} else {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Permission denied"})
	}
	c.Abort()
}

func (s *Server) handleGetIngredientTable(c *gin.Context) {
	filename := "ingredients.csv"
	filePath := filepath.Join("assets", filename)

	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	// Set headers for CSV file download
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "text/csv")

	// Copy the file contents to the response writer
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}
}

func (s *Server) handleHealthCheck(c *gin.Context) {
	log.Println("Health check received")
	c.Status(http.StatusOK)
	c.String(http.StatusOK, "OK")

}
