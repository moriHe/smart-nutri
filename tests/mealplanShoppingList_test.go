package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetMealplanItemsShoppingListSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/mealplan/shopping-list", nil)
	r.ServeHTTP(w, req)
	response := `{"data":[{"id":1,"mealplanItem":{"id":1,"recipeName":"Spaghetti","date":"2023-08-22","portions":2,"meal":"DINNER"},` +
		`"recipeIngredient":{"id":1,"name":"Tomaten","amountPerPortion":100,"unit":"GRAM","market":"Rewe","isBio":true}},{"id":2,"mealplanItem"` +
		`:{"id":2,"recipeName":"Pizza","date":"2023-08-22","portions":1,"meal":"BREAKFAST"},"recipeIngredient":{"id":2,"name":"Knoblauch",` +
		`"amountPerPortion":200,"unit":"GRAM","market":"Rewe","isBio":false}}]}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, w.Body.String())
}
