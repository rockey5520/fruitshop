package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary Show details of all fruits
// @Description Get details of all fruits
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Fruit
// @Router /server/api/v1/fruits [get]
// FindFruit returns fruit details if fruit exists in the store
func FindFruit(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var fruit models.Fruit

	if err := db.Where("name = ?", c.Param("name")).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fruit})
}

// @Summary Show details of a fruit item
// @Description Get details of a fruit item
// @Accept  json
// @Produce  json
// @Param name path string true "Fruit identifier"
// @Success 200 {object} models.Fruit
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/fruits/{name} [get]
// FindFruits will retuen all fruits exists within the fruitshop
func FindFruits(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var fruits []models.Fruit
	db.Find(&fruits)

	c.JSON(http.StatusOK, gin.H{"data": fruits})
}
