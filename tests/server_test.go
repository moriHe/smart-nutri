package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/moriHe/smart-nutri/api"
)

func TestGetAllRecipesSucc(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes", nil)
	router.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":[{"id":1,"name":"Spaghetti"},{"id":2,"name":"Pizza"}]}`, w.Body.String())
}

func TestGetRecipeByIdSucc(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/1", nil)
	router.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":{"id":1,"name":"Spaghetti","ingredients":[{"id":1,"ingredientId":1,"name":"Tomaten"},{"id":2,"ingredientId":2,"name":"Knoblauch"}]}}`, w.Body.String())

}

func TestGetRecipeByIdBadReq(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/3", nil)
	router.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Bad Request: No recipe found with id 3"}}`, w.Body.String())
}

func TestPostRecipeNoIngredientsSucc(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer([]byte(`{
		"name": "Wantan"
	}`)))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added recipe"}`, w.Body.String())
}

func TestPostRecipeIngredientsSucc(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer([]byte(`{
		"name": "Wantan",
		"ingredients": [1,2]
	}`)))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added recipe"}`, w.Body.String())
}
