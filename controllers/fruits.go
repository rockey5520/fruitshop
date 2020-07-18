package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindFruit returns fruit details if fruit exists in the store
func FindFruit(c *gin.Context) {
	var fruit models.Fruit

	if err := models.DB.Where("id = ?", c.Param("id")).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fruit})
}

// FindFruits will retuen all fruits exists within the fruitshop
func FindFruits(c *gin.Context) {

	var fruits []models.Fruit
	models.DB.Find(&fruits)

	c.JSON(http.StatusOK, gin.H{"data": fruits})
}
