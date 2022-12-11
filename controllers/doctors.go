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

//post doctors
func CreateDoctors(c *gin.Context){
	var input models.CreateDoctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create doctor
	doctor := models.Doctor{Name: input.Name, Designation: input.Designation, Ratings: input.Ratings}
	models.DB.Create(&doctor)

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}
