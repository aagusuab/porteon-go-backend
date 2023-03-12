package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"porteonBackend/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/truck") // new

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
