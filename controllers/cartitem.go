package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddItemInCart will add users choosen fruits to the cart list
func AddItemInCart(c *gin.Context) {
	// Bind the input payload to schema for validations
	var input models.CartItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Fetch customer information
	var customer models.Customer
	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}
	// Fetch cart information
	var cart models.Cart
	if err := models.DB.Where("customer_id = ?", customer.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found!"})
		return
	}
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: cart.ID, Name: input.Name}
	if input.Count > 0 {
		// Create/update fruit to the cart

		cartItem.Count = input.Count
		if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&cartItem) // create new record from newUser
			}
		} else {
			models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).Update("count", input.Count)
		}
	} else if input.Count == 0 {
		models.DB.Where("cart_id = ? AND name = ?", cart.ID, input.Name).Delete(&cartItem)
		//if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {

	}
	fruit := models.Fruit{Name: input.Name}
	// recalcuate the payment for the cart
	RecalcualtePayments(customer, cart, cartItem, fruit)

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(customer models.Customer, cart models.Cart, cartItem models.CartItem, fruit models.Fruit) {
	// Recalcualte the payments
	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	for _, item := range cartItems {
		models.DB.First(&fruit)
		totalCost += float64(item.Count) * fruit.Price
	}
	models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", totalCost)

	if fruit.Name == "Apple" && cartItem.Count >= 7 {
		var restCost float64
		var appleCost float64
		for _, item := range cartItems {
			if item.Name != "Apple" {
				var fruit models.Fruit
				models.DB.Where("name = ?", item.Name).First(&fruit)
				restCost += float64(item.Count) * fruit.Price
			} else {
				var fruit models.Fruit
				models.DB.Where("name = ?", item.Name).First(&fruit)
				appleCost += float64(item.Count) * fruit.Price
			}
			discount := ((float64(item.Count) * fruit.Price) / 100) * 10
			totalCost = (appleCost - discount) + restCost
		}
		models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", totalCost)
	}

}
