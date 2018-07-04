package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": message()})
	})

	engine.Run(port())
}

func message() string {
	msg := os.Getenv("MESSAGE")
	if len(msg) == 0 {
		msg = "Hello Cloud Native Night Muc with SquareScale!"
	}
	return msg
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	return ":" + port
}
