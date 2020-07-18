package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCartList will retuen all of the cart list for a given customer
func GetAllCartItems(c *gin.Context) {

	var cartItems []models.CartItem
	customer := models.Customer{}
	cart := models.Cart{}

	if err := models.DB.Where("login_id = ?", c.Param("login_id")).Find(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	if err := models.DB.Where("customer_id = ?", customer.ID).Find(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found!"})
		return
	}
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CusCartItems record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

// FindCart will fetch the details about the cart of the customer
func FindCart(c *gin.Context) {
	customer := models.Customer{}
	cart := models.Cart{}

	models.DB.Where("login_id = ?", c.Param("login_id")).Find(&customer)

	models.DB.Where("customer_id = ?", customer.ID).Find(&cart)
	fmt.Println(cart)
	c.JSON(http.StatusOK, gin.H{"data": cart})

}

func CreateCart(customer models.Customer) {
	models.DB.Where("customer_login_id = ?", customer.LoginId).Find(&customer)

	cart := models.Cart{CustomerId: customer.ID, Total: 0.0}

	if err := models.DB.Create(&cart).Error; err != nil {
		panic("Unable to create cart")
	}

}
