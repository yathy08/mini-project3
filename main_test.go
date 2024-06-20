// tests/main_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/gock.v1"
	"github.com/yourusername/yourproject/internal/domain"
	"github.com/yourusername/yourproject/internal/handler"
)

func TestGetAll(t *testing.T) {
	gock.New("https://reqres.in").
		Get("/api/users").
		Reply(200).
		JSON(map[string]interface{}{
			"data": []domain.User{
				{ID: 1, Email: "garzao@e.o.cara"},
			},
		})

	router := gin.Default()
	router.GET("/", handler.GetAll)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %v; got %v", http.StatusOK, rr.Code)
	}

	var res handler.Users
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

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
	router := gin.Default()
	router.GET("/:id", handler.GetByID)

	req, _ := http.NewRequest("GET", "/invalid", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected %v; got %v", http.StatusBadRequest, rr.Code)
	}
}
