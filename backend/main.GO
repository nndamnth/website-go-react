package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// API backend
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Golang Backend 🚀",
		})
	})

	// Serve React build
	r.Static("/assets", "../frontend/dist/assets")
	r.StaticFile("/", "../frontend/dist/index.html")

	// Kalau buka route lain (misal /about), tetap arahkan ke React
	r.NoRoute(func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})

	r.Run(":8080") // buka di http://localhost:8080
}
