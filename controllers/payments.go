package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Pay method takes the payment and resets cart, cartitems, coupons, discounts
func Pay(c *gin.Context) {

	// Validate input
	var payment PayInput
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get customer
	var customer models.Customer
	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	// Get cart
	cart := models.Cart{}
	if err := models.DB.Where("customer_id = ?", customer.ID).Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found!"})
		return
	}
	if cart.Total == payment.Amount {
		// Empyt cart items table
		var cartItems []models.CartItem
		models.DB.Find(&cartItems)
		for _, item := range cartItems {
			models.DB.Where("cart_id = ? AND name = ?", cart.ID, item.Name).Delete(&item)
		}
		// Set Cart amoun to 0
		models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", 0)
		//update orangecoupon to NOTAPPLIED
		var coupon models.Coupon
		models.DB.Model(&coupon).Where("cart_id = ? and name = ?", cart.ID, "ORANGE30").Update("status", "NOTAPPLIED")
		// discounts table reset to default values
		var discount models.Discount
		models.DB.Model(&discount).Where("customer_id = ?", customer.ID).Update("status", "NOTAPPLIED")
		var pay models.Payment
		models.DB.Model(&pay).Where("cart_id = ?", cart.ID).Update("status", "PAID")

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment amount mismatched with the cart total"})
		return
	}
	payment.Status = "PAID"

	c.JSON(http.StatusOK, gin.H{"data": payment})
}
