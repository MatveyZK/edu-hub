package main

import (
	"github.com/MatveyZK/edu-hub/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", handlers.Health)
	r.Run(":8080")
}
