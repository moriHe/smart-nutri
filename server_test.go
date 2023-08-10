package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/moriHe/smart-nutri/api"
)

func TestHandleGetAllRecipes(t *testing.T) {
	router := api.StartGinServer(Db, "localhost:5432")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/recipes", nil)
	router.R.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
