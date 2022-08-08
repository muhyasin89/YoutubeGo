package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello this is init gin gorm"})
	})

	r.Run(":8080")

}
