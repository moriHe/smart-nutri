package tests

import (
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

func TestGetMeaplanItemSucc(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mealPlan/item/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"data":{"id":"1","date":"2023-08-22","meal":"DINNER","portions":2,"recipe":{"id":1,"name":"Spaghetti","recipeIngredients":[{"id":1,"name":"Tomaten","amountPerPortion":100,"unit":"GRAM","market":"Rewe","isBio":true},{"id":2,"name":"Knoblauch","amountPerPortion":200,"unit":"GRAM","market":"Rewe","isBio":false}]}}}`, w.Body.String())
}

func TestGetMeaplanItemBadReq(t *testing.T) {
	r := startServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mealPlan/item/1000", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":{"status":400,"message":"Bad Request: no rows in result set"}}`, w.Body.String())
}
