package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetMeaplanSuccHasMealplans(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/mealplan/2023-08-22", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":[{"id":1,"recipeName":"Spaghetti","date":"2023-08-22","meal":"DINNER"},{"id":2,"recipeName":"Pizza","date":"2023-08-22","meal":"BREAKFAST"}]}`, w.Body.String())
}

func TestGetMealplanSuccNoMeaplans(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/mealplan/2100-08-22", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":[]}`, w.Body.String())
}

func TestGetMealplanBadReqInvalidDate(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/familys/1/mealplan/INVALID_DATE", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Invalid Date. Use format YYYY-MM-DD"}}`, w.Body.String())
}

func TestGetMealplanItemSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mealplan/item/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":{"id":"1","date":"2023-08-22","meal":"DINNER","portions":2,"recipe":{"id":1,"name":"Spaghetti","recipeIngredients":[{"id":1,"name":"Tomaten","amountPerPortion":100,"unit":"GRAM","market":"Rewe","isBio":true},{"id":2,"name":"Knoblauch","amountPerPortion":200,"unit":"GRAM","market":"Rewe","isBio":false}]}}}`, w.Body.String())
}

func TestGetMealplanItemBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mealplan/item/1000", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Bad Request: no rows in result set"}}`, w.Body.String())
}

// TODO Check if item is posted correctly to db
func TestPostMealplanItemSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1/mealplan", bytes.NewBuffer([]byte(`{
		"recipeId": 2,
		"date": "2023-08-22",
		"meal": "DINNER",
		"portions": 2
	}`)))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Added mealplan item"}`, w.Body.String())
}

// TODO: Test other missing payloads
// TODO: Add proper error messages
func TestPostMealplanItemBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/familys/1/mealplan", bytes.NewBuffer([]byte(`{
		"date": "2023-08-22",
		"meal": "DINNER",
		"portions": 2
	}`)))
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Step 2: Failed to create mealplan item: ERROR: insert or update on table \"mealplans\" violates foreign key constraint \"mealplans_recipe_id_fkey\" (SQLSTATE 23503)"}}`, w.Body.String())
}

func TestDeleteMealplanItemSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/mealplan/item/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":"Deleted mealplan item"}`, w.Body.String())
}

func TestDeleteMealplanItemBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/mealplan/item/1000", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Mealplan item does not exist"}}`, w.Body.String())
}
