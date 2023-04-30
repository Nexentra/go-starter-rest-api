package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	// create a new Gin router
	router := gin.Default()

	// register the "ping" route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// create a new HTTP request to the "ping" route
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new HTTP recorder to capture the response
	w := httptest.NewRecorder()

	// dispatch the HTTP request to the router
	router.ServeHTTP(w, req)

	// check that the HTTP status code is 200 OK
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check that the response body matches what we expect
	expected := `{"message":"pong"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
