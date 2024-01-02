package api

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/middleware"
	"github.com/moriHe/smart-nutri/storage"
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
	router.Use(middleware.AuthMiddleware(store, authClient))
	server.recipeRoutes(router)
	server.mealPlanRoutes(router)
	server.mealplanShoppingListRoutes(router)

	router.Run(url)

	return router, nil

}
