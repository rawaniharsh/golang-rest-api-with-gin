package controllers

import (
	"goRest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//get all doctors
func FindDoctors(c *gin.Context) {
	var doctors []models.Doctor
	models.DB.Find(&doctors)

	c.JSON(http.StatusOK, gin.H{"data": doctors})
}
