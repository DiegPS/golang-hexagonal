package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET(
		"/", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"message": "Welcome, this is a basic gin (https://github.com/gin-gonic/gin) server deployed on Zeabur (https://zeabur.com)",
				},
			)
		},
	)

	panic(r.Run(":" + port))
}
