// internal/handler/handler.go

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/repository"
	"github.com/yathy08/mini-project3/internal/service"
)

const apiURL = "https://reqres.in/api/users"

// GetAll handles the request to get all users.
func GetAll(c *gin.Context) {
	data, err := repository.FetchUsers(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	users, err := service.UnmarshalUsers(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetByID handles the request to get a user by ID.
func GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	data, err := repository.FetchUsers(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	users, err := service.UnmarshalUsers(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	user := service.FilterByID(users, id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
