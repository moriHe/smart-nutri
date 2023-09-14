package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetMealplanItemsShoppingListSuccHasItems(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/mealplan/shopping-list", nil)
	r.ServeHTTP(w, req)
	response := `{"data":[{"id":1,"market":"REWE","isBio":true,"mealplanItem":{"id":1,"recipeName":"Spaghetti","date":"2023-08-22","` +
		`portions":2,"meal":"DINNER"},"recipeIngredient":{"id":1,"name":"Tomaten","amountPerPortion":100,"unit":"GRAM"}},{"id":2,"market":` +
		`"EDEKA","isBio":false,"mealplanItem":{"id":2,"recipeName":"Pizza","date":"2023-08-22","portions":1,"meal":"BREAKFAST"},"recipeIngredient":` +
		`{"id":2,"name":"Knoblauch","amountPerPortion":200,"unit":"GRAM"}}]}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, w.Body.String())
}

func TestGetMealplanItemsShoppingListSuccNoItems(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/2/mealplan/shopping-list", nil)
	r.ServeHTTP(w, req)
	response := `{"data":[]}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, w.Body.String())
}

func TestPostMealplanItemShoppingListSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1/mealplan/1/shopping-list", bytes.NewBuffer([]byte(`{
		"market": "REWE",
		"isBio": true,
		"recipeIngredientId": 1
	}`)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added mealplan item to shopping list"}`, w.Body.String())
}

func TestPostMealplanItemShoppingListBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1000/mealplan/1000/shopping-list", bytes.NewBuffer([]byte(`{
		"market": "REWE",
		"isBio": true,
		"recipeIngredientId": 1000
	}`)))
	r.ServeHTTP(w, req)
	response := `{"error":{"status":400,"message":"Error: Failed to post mealplan item shopping list: ERROR: insert or ` +
		`update on table \"mealplans_shopping_list\" violates foreign key constraint \"mealplans_shopping_list_family_id_fkey\" (SQLSTATE 23503)"}}`
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, response, w.Body.String())
}

func TestDeleteMealplanItemShoppingListSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/mealplan/shopping-list/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Deleted shopping list item"}`, w.Body.String())
}

func TestDeleteMealplanItemShoppingListBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/mealplan/shopping-list/1000", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Shopping list item does not exist"}}`, w.Body.String())
}
