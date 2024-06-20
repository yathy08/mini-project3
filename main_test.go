// tests/main_test.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/domain"
	"github.com/yathy08/mini-project3/internal/handler"
	"gopkg.in/h2non/gock.v1"
)

// Example adjustment for TestGetAll
func TestGetAll(t *testing.T) {
	// Mock API response
	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"data": []domain.User{
				{ID: 1, Email: "garzao@e.o.cara"},
			},
		})

	// Setup Gin router
	router := gin.Default()
	router.GET("/", handler.GetAll)

	// Perform GET request
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert status code
	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	// Unmarshal response body
	var res handler.Users
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// Assert response body
	expected := []domain.User{{ID: 1, Email: "garzao@e.o.cara"}}
	if !reflect.DeepEqual(res.Data, expected) {
		t.Fatalf("expected %v; got %v", expected, res.Data)
	}
}

func TestGetByID(t *testing.T) {
	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"data": []domain.User{
				{ID: 1, Email: "garzao@e.o.cara"},
			},
		})

	router := gin.Default()
	router.GET("/:id", handler.GetByID)

	req, _ := http.NewRequest("GET", "/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	var res domain.User
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expected := domain.User{ID: 1, Email: "garzao@e.o.cara"}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v; got %v", expected, res)
	}
}

func TestGetByIDInvalid(t *testing.T) {
	// Initialize Gin router
	router := gin.Default()

	// Register the GetByID handler
	router.GET("/:id", handler.GetByID)

	// Custom debug endpoint to list registered routes (optional for testing)
	router.GET("/debug/routes", func(c *gin.Context) {
		// Collect registered routes
		routes := make(map[string][]string)
		for _, route := range router.Routes() {
			routes[route.Method] = append(routes[route.Method], route.Path)
		}
		c.JSON(http.StatusOK, gin.H{"routes": routes})
	})

	// Create a GET request with an invalid ID
	req, err := http.NewRequest("GET", "/invalid", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Serve the request and record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert the response status code should be BadRequest (400)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %v; got %v", http.StatusBadRequest, rr.Code)
	}

	// Log registered routes for debugging
	log.Println("Registered Routes:", router.Routes())
}