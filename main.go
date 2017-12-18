package main

import (
  "os"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.String(200, "API")
  })
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
  port := os.Getenv("PORT")
	r.Run(port)
}
