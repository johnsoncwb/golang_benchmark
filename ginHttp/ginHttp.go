package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/:name", GinHelloWorld)
	r.Run("localhost:9000")
}

func GinHelloWorld(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Olá %s!!!", name),
	})
}
