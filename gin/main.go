package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world222"})
}

func main() {
	engine := gin.Default()

	engine.GET("/", handler)
	engine.Run(":3000")
}
