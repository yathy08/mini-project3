// tests/main_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/domain"
	"github.com/yathy08/mini-project3/internal/handler"
	"gopkg.in/h2non/gock.v1"
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

	t.Logf("Response Body: %s", rr.Body.String())
	t.Logf("Unmarshaled Data: %+v", res)

	expected := []domain.User{{ID: 1, Email: "garzao@e.o.cara"}}
	if len(res.Data) != len(expected) {
		t.Fatalf("expected %d items; got %d items", len(expected), len(res.Data))
	}

	for i := range res.Data {
		if res.Data[i].ID != expected[i].ID || res.Data[i].Email != expected[i].Email {
			t.Fatalf("expected %v; got %v", expected[i], res.Data[i])
		}
	}
}

func TestGetByID(t *testing.T) {
	// Mock API response
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
	if res.ID != expected.ID || res.Email != expected.Email {
		t.Fatalf("expected %v; got %v", expected, res)
	}
}
