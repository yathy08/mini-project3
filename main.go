// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/handler"
)

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/", handler.GetAll)
	r.GET("/:id", handler.GetByID)

	log.Println("Server is running on http://localhost:3000")
	r.Run(":3000")
}
