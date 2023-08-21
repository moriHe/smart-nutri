package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert"
	"github.com/moriHe/smart-nutri/api"
)

func startServer() *api.Server {
	r := api.StartGinServer(Db, os.Getenv("DOCKER_TEST_SERVER_URL"))
	return r
}

func TestGetAllRecipesSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes", nil)
	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":[{"id":1,"name":"Spaghetti"},{"id":2,"name":"Pizza"}]}`, w.Body.String())
}

func TestGetRecipeByIdSucc(t *testing.T) {
	r := startServer()
	// TODO Return all fields
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/1", nil)
	r.R.ServeHTTP(w, req)
	res := `{"data":{"id":1,"name":"Spaghetti","recipeIngredients":` +
		`[{"id":1,"name":"Tomaten","amountPerPortion":"100","unit":"GRAM","market":` +
		`"Rewe","isBio":true},{"id":2,"name":"Knoblauch","amountPerPortion":"200",` +
		`"unit":"GRAM","market":"Rewe","isBio":false}]}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res, w.Body.String())

}

func TestGetRecipeByIdBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/1000", nil)
	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Bad Request: No recipe found with id 1000"}}`, w.Body.String())
}

func TestPostRecipeSuccNoIngredients(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer([]byte(`{
		"name": "Wantan"
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added recipe"}`, w.Body.String())
}

func TestPostRecipeSuccIngredients(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	body := `{"name": "Wantan","recipeIngredients": [{"ingredientId": 1,` +
		`"amountPerPortion": 100,"unit": 1,"market": 1,"isBio": true},{"ingredientId": 2,` +
		`"amountPerPortion": 200,"unit": 2,"market": 2,"isBio": false}]}`
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer([]byte(body)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO: Return payload in response
	assert.Equal(t, `{"data":"Added recipe"}`, w.Body.String())
}

func TestPostRecipeBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer([]byte(`{
		"hello": "world"
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"No recipe name specified"}}`, w.Body.String())
}

func TestPostRecipeIngredientSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes/2/ingredients", bytes.NewBuffer([]byte(`{
		"ingredientId": 2,
		"amount": 1,
		"unit": 1,
		"market": 1,
		"isBio": true
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added recipe ingredient"}`, w.Body.String())
}

func TestPostRecipeIngredientBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/recipes/1000/ingredients", bytes.NewBuffer([]byte(`{
		"ingredientId": 2
	}`)))

	r.R.ServeHTTP(w, req)
	// TODO: Reset db after each test
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Failed to post recipe_ingredient: ERROR: insert or update on table \"recipes_ingredients\" violates foreign key constraint \"recipes_ingredients_recipe_id_fkey\" (SQLSTATE 23503)"}}`, w.Body.String())
}

func TestPatchRecipeNameSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/2", bytes.NewBuffer([]byte(`{
		"name": "Beyond Burger"
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe name updated"}`, w.Body.String())
}

func TestPatchRecipeNameBadReqNoName(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/2", bytes.NewBuffer([]byte(`{
		"hello": "world"
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"No recipe name specified"}}`, w.Body.String())
}

func TestPatchRecipeNameBadReqNoId(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/1000", bytes.NewBuffer([]byte(`{
		"name": "Pasta"
	}`)))

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe does not exist"}}`, w.Body.String())
}

func TestDeleteRecipeSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/2", nil)

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe deleted"}`, w.Body.String())
}

func TestDeleteRecipeBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/1000", nil)

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe does not exist"}}`, w.Body.String())
}

func TestDeleteRecipeIngredientSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/ingredients/1", nil)

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe ingredient deleted"}`, w.Body.String())
}
func TestDeleteRecipeIngredientBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/ingredients/1000", nil)

	r.R.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe ingredient does not exist"}}`, w.Body.String())
}
