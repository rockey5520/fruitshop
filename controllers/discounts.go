package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindDiscounts will return all discounts with status as APPLIED available within the fruitshop
func FindDiscounts(c *gin.Context) {

	customer := models.Customer{
		LoginId: c.Param("login_id"),
	}
	var discounts []models.Discount

	//models.DB.First(&customer)
	// Where("customer_login_id = ?", c.Param("login_id")).

	//models.DB.First(&customer, c.Param("login_id"))
	models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).
		Preload("Discounts").
		Find(&customer)
	discounts = customer.Discounts

	c.JSON(http.StatusOK, gin.H{"data": discounts})
}
