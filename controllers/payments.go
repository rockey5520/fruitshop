package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var payment PayInput
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get cart
	cart := models.Cart{}
	if err := db.Where("ID = ? AND status = ?", payment.CartID, "OPEN").Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found or Payment is made on already paid cart!"})
		return
	}

	if cart.Total == payment.Amount && cart.Total != 0 && payment.Amount > 0 {
		// Empyt cart items table
		var cartItems []models.CartItem
		db.Find(&cartItems)

		// Set Cart amount to 0
		db.Model(&cart).Where("ID = ?", payment.CartID).Update("total", 0).Update("status", "CLOSED")

		pay := models.Payment{
			CartId: payment.CartID,
			Amount: payment.Amount,
			Status: "PAID",
		}

		db.Create(&pay)

		newCart := models.Cart{
			CustomerId: payment.CustomerID,
			Total:      0.0,
			Status:     "OPEN",
		}
		db.Create(&newCart)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment amount mismatched with the cart total"})
		return
	}

	var customer models.Customer
	db.Where("ID = ?", payment.CustomerID)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
