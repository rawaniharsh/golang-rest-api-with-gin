package main

import (
	"goRest/controllers"
	"goRest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Data": "Hello World"})
	})

	r.GET("/doctors", controllers.FindDoctors)

	r.POST("/doctors", controllers.CreateDoctors)

	models.ConnectDatabase()

	r.Run()
}
