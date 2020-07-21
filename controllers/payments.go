package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Payment endpoint
// @Description This end point will update payment details of cart into the database
// @Accept  json
// @Produce  json
// @Param Input body models.Payment true "Payment input request"
// @Param cart_id path string true "Customer identifier"
// @Success 200 {object} models.Payment
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/pay/{cart_id} [post]
// Pay method takes the payment and resets cart, cartitems, coupons, discounts
func Pay(c *gin.Context) {

	// Validate input
	var payment PayInput
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get cart
	cart := models.Cart{}
	if err := models.DB.Where("ID = ? AND status = ?", payment.CartID, "NOTPAID").Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found or Payment is made on already paid cart!"})
		return
	}

	if cart.Total == payment.Amount && cart.Total != 0 && payment.Amount > 0 {
		// Empyt cart items table
		var cartItems []models.CartItem
		models.DB.Find(&cartItems)

		// Set Cart amoun to 0
		models.DB.Model(&cart).Where("ID = ?", payment.CartID).Update("total", 0).Update("status", "CLOSED")

		var pay models.Payment
		models.DB.Model(&pay).Where("cart_id = ?", cart.ID).Update("status", "PAID")
		newCart := models.Cart{
			CustomerId: payment.CustomerID,
			Total:      0.0,
			Status:     "OPEN",
		}
		models.DB.Create(&newCart)
		var cart models.Cart

		models.DB.Where("status = ?", "CLOSED").Find(&cart)
		newPayment := models.Payment{
			CartId: cart.ID,
			Amount: 0.0,
			Status: "NOTPAID",
		}
		models.DB.Create(&newPayment)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment amount mismatched with the cart total"})
		return
	}
	payment.Status = "PAID"
	var customer models.Customer
	models.DB.Where("ID = ?", payment.CustomerID)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
