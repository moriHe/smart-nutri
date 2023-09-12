package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetAllRecipesSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/recipes", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":[{"id":1,"name":"Spaghetti"}]}`, w.Body.String())
}

func TestGetRecipeByIdSucc(t *testing.T) {
	r := startServer()
	// TODO Return all fields
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/1", nil)
	r.ServeHTTP(w, req)
	res := `{"data":{"id":1,"name":"Spaghetti","defaultPortions":1,"defaultMeal":"BREAKFAST","recipeIngredients":` +
		`[{"id":1,"name":"Tomaten","amountPerPortion":100,"unit":"GRAM","market":` +
		`"Rewe","isBio":true},{"id":2,"name":"Knoblauch","amountPerPortion":200,` +
		`"unit":"GRAM","market":"Rewe","isBio":false}]}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res, w.Body.String())

}

func TestGetRecipeByIdBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes/1000", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Bad Request: No recipe found with id 1000"}}`, w.Body.String())
}

func TestPostRecipeSuccNoIngredients(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1/recipes", bytes.NewBuffer([]byte(`{
		"name": "Wantan",
		"defaultPortions": 1.5,
		"defaultMeal": "DINNER"
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":{"id":3}}`, w.Body.String())
}

func TestPostRecipeSuccIngredients(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	body := `{"name": "Wantan","defaultPortions":1.5,"defaultMeal":"DINNER","recipeIngredients": [{"ingredientId": 1,` +
		`"amountPerPortion": 100,"unit": "GRAM","marketId": 1,"isBio": true},{"ingredientId": 2,` +
		`"amountPerPortion": 200,"unit": "MILLILITER","marketId": 2,"isBio": false}]}`
	req, _ := http.NewRequest("POST", "/familys/1/recipes", bytes.NewBuffer([]byte(body)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO: Return payload in response
	assert.Equal(t, `{"data":{"id":4}}`, w.Body.String())
}

// TODO: Add bad request tests for missing ingriedientId, amountPerPortion etc
// TODO: Also add test family id not correct
func TestPostRecipeBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1/recipes", bytes.NewBuffer([]byte(`{
		"hello": "world",
		"defaultPortions": 1.5,
		"defaultMeal": "DINNER"
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"No recipe name specified"}}`, w.Body.String())
}

func TestPostRecipeIngredientSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/recipes/2/recipeingredient", bytes.NewBuffer([]byte(`{
		"ingredientId": 2,
		"amount": 1,
		"unit": "GRAM",
		"marketId": 1,
		"isBio": true
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added recipe ingredient"}`, w.Body.String())
}

func TestPostRecipeIngredientBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/recipes/1000/recipeingredient", bytes.NewBuffer([]byte(`{
		"ingredientId": 2,
		"unit": "GRAM"
	}`)))

	r.ServeHTTP(w, req)
	// TODO: Reset db after each test
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Step 2: Failed to post recipe_ingredient: ERROR: insert or update on table \"recipes_ingredients\" violates foreign key constraint \"recipes_ingredients_recipe_id_fkey\" (SQLSTATE 23503)"}}`, w.Body.String())
}

func TestPatchRecipeNameSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/2", bytes.NewBuffer([]byte(`{
		"name": "Beyond Burger"
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe name updated"}`, w.Body.String())
}

func TestPatchRecipeNameBadReqNoName(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/2", bytes.NewBuffer([]byte(`{
		"hello": "world"
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"No recipe name specified"}}`, w.Body.String())
}

func TestPatchRecipeNameBadReqNoId(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/recipes/1000", bytes.NewBuffer([]byte(`{
		"name": "Pasta"
	}`)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe does not exist"}}`, w.Body.String())
}

func TestDeleteRecipeSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/2", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe deleted"}`, w.Body.String())
}

func TestDeleteRecipeBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/1000", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe does not exist"}}`, w.Body.String())
}

func TestDeleteRecipeIngredientSucc(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/recipeingredient/1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Recipe ingredient deleted"}`, w.Body.String())
}
func TestDeleteRecipeIngredientBadReq(t *testing.T) {
	r := startServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/recipes/recipeingredient/1000", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Recipe ingredient does not exist"}}`, w.Body.String())
}
