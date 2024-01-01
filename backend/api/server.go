package api

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
	"google.golang.org/api/option"
)

type Server struct {
	store storage.Storage
	Auth  *auth.Client
}

func StartGinServer(store storage.Storage, url string) (*gin.Engine, error) {
	router := gin.Default()
	opt := option.WithCredentialsFile("/Users/moritzhettich/prv/smart-nutri/backend/firebase-private-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	server := &Server{store: store, Auth: authClient}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update with your Angular app's origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	// Use the CORS middleware
	router.Use(cors.New(config))

	server.userRoutes(router)
	server.recipeRoutes(router)
	server.mealPlanRoutes(router)
	server.mealplanShoppingListRoutes(router)

	router.Run(url)

	return router, nil

}

func errorResponse(c *gin.Context, err error) bool {
	if err != nil {
		if requestErr, ok := err.(*types.RequestError); ok {
			c.JSON(requestErr.Status, gin.H{"error": requestErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return true
	}

	return false
}

func handleResponse[T any](c *gin.Context, successResponse T, err error) {
	if err != nil {
		if requestErr, ok := err.(*types.RequestError); ok {
			c.JSON(requestErr.Status, gin.H{"error": requestErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": successResponse})
}
