package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindCart will fetch the details about the cart of the customer

func FindCart(c *gin.Context) {
	customer := models.Customer{}
	var cart models.Cart

	models.DB.Where("login_id = ?", c.Param("login_id")).Find(&customer)
	models.DB.
		Preload("Cart").
		Find(&customer)

	cart = customer.Cart
	cart.ID = customer.ID
	c.JSON(http.StatusOK, gin.H{"data": cart})

}

func CreateCart(customer models.Customer) {
	models.DB.Where("customer_login_id = ?", customer.LoginId).Find(&customer)

	cart := models.Cart{CustomerId: customer.ID, Total: 0.0}

	if err := models.DB.Create(&cart).Error; err != nil {
		panic("Unable to create cart")
	}

}
