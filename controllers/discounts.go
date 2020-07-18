package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindDiscounts will return all discounts with status as APPLIED available within the fruitshop
func FindDiscounts(c *gin.Context) {

	customer := models.Customer{
		CustomerLoginId: c.Param("login_id"),
	}
	var discounts models.Discount

	models.DB.First(&customer)
	discounts.CustomerLoginId = customer.CustomerLoginId

	models.DB.Where("customer_login_id = ?", c.Param("login_id")).
		Preload("Discounts").
		Find(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
