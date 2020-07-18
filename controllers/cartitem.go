package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddItemInCart will add users choosen fruits to the cart list
func AddItemInCart(c *gin.Context) {
	var input models.CartItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	var cart models.Cart
	if err := models.DB.Where("customer_id = ?", customer.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found!"})
		return
	}
	cartItem := models.CartItem{CartID: cart.ID, Name: input.Name}
	if input.Count > 0 {
		// Create/update fruit to the cart
		cartItem.Count = input.Count
		if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&cartItem) // create new record from newUser
			}
		} else {
			models.DB.Model(&cartItem).Where("ID = ? and name = ?", cart.ID, input.Name).Update("count", input.Count)
		}
	} else if input.Count == 0 {
		models.DB.Unscoped().Delete(&cartItem)
		//if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {

	}

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}
