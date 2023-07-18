package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Recipe struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
}

var db *pgx.Conn

func getAllRecipes(c *gin.Context) {
	rows, err := db.Query(context.Background(), "select * from recipes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return
	}

	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe

		err = rows.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return
		}
		recipes = append(recipes, recipe)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
	}
	fmt.Print(recipes)
	c.JSON(http.StatusOK, recipes)

}

func getRecipeById(c *gin.Context) {
	c.Param("id")
}

func postRecipe(c *gin.Context) {

}

func patchRecipe(c *gin.Context) {

}

func deleteRecipe(c *gin.Context) {

}

func main() {
	var dbErr error
	db, dbErr = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", dbErr)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	var name string
	dbErr = db.QueryRow(context.Background(), "select name from ingredients where id=$1", 42).Scan(&name)
	if dbErr != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", dbErr)
		os.Exit(1)
	}

	fmt.Println(name)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello World!")
	})

	router.GET("/recipes", getAllRecipes)
	router.GET("/recipes/:id", getRecipeById)
	router.POST("/recipes", postRecipe)
	router.PATCH("/recieps", patchRecipe)
	router.DELETE("/recipes/:id", deleteRecipe)

	router.Run("localhost:8080")
}
