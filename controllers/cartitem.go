package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CartItemInput struct {
	CartId uint   `json:"cartid" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Count  int    `json:"count"`
}

// @Summary Show details of a all items in a cart
// @Description Get details of contents of the cart
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"
// @Success 200 {object} models.CartItem
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/cartitem/{cart_id} [get]
// GetAllCartList will retuen all of the cart list for a given customer
func GetAllCartItems(c *gin.Context) {
	s := make([]models.CartItemResponse, 0)

	var cartItems []models.CartItem
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&cartItems)

	for _, cartItem := range cartItems {
		var fruit models.Fruit
		models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		s = append(s, models.CartItemResponse{
			Name:        fruit.Name,
			CostPerItem: fruit.Price,
			Count:       cartItem.Quantity,
			ItemTotal:   cartItem.ItemTotal,
		})

	}
	c.JSON(http.StatusOK, gin.H{"data": s})
}

// @Summary Creates/Updated item in the cart
// @Description This end point will record cart item details into the database
// @Accept  json
// @Produce  json
// @Param Input body models.CartItem true "Input request"
// @Param login_id path string true "Customer identifier"
// @Success 200 {object} models.CartItem
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/cartitem/{login_id} [post]
// CreateUpdateItemInCart will add users choosen fruits to the cart list
func CreateUpdateItemInCart(c *gin.Context) {
	// Bind the input payload to schema for validations
	var input CartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fruit models.Fruit
	if err := models.DB.Where("name = ?", input.Name).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	fmt.Println("fruit id ", fruit.ID)
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: input.CartId, FruitID: fruit.ID, ItemTotal: fruit.Price * float64(input.Count), ItemDiscountedTotal: 0.0}
	if input.Count > 0 {
		// Create/update fruit to the cart
		cartItem.Quantity = input.Count
		if err := models.DB.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", input.CartId, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&cartItem) // create new record from newUser
			}
		} else {
			models.DB.Model(&cartItem).Where("cart_id = ?  AND fruit_id = ? ", input.CartId, fruit.ID).
				Update("quantity", input.Count).
				Update("fruit_id", fruit.ID).
				Update("item_total", float64(input.Count)*fruit.Price).
				Update("item_discounted_total", 0.0)
		}
	} else if input.Count == 0 {
		models.DB.Where("cart_id = ? AND fruit_id = ?", input.CartId, fruit.ID).Delete(&cartItem)

	}

	// RecalcuateItem payment for the item in the cart
	ApplySingleItemDiscounts(cartItem)
	ApplyDualItemDiscounts(cartItem)
	// Recalcuate the payment for the cart
	RecalcualtePayments(cartItem.CartID)

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(cartID uint) {
	// Recalcualte the payments

	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	var totalDiscountedCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
		totalDiscountedCost += item.ItemDiscountedTotal
	}
	var cart models.Cart
	if err := models.DB.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}
	models.DB.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
	/* var payment models.Payment
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&payment).Error; err != nil {
		fmt.Println("Error ", err)
	}
	models.DB.Model(&payment).Where("cart_id = ?", cart.ID).Update("amount", totalCost) */

}
