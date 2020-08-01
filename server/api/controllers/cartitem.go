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
	db := c.MustGet("db").(*gorm.DB)
	var cartItems []models.CartItem
	db.Where("cart_id = ?", c.Param("cart_id")).Find(&cartItems)

	for _, cartItem := range cartItems {
		var fruit models.Fruit
		db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
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
func (server *Server) CreateUpdateItemInCart(w http.ResponseWriter, r *http.Request) {

	cartItemCreated, err := cartItem.SaveUpdateCartItem(server.DB)
	// RecalcuateItem payment for the item in the cart
	ApplySingleItemDiscounts(cartItemCreated, c)
	ApplyDualItemDiscounts(cartItemCreated, c)
	// Recalcuate the payment for the cart
	RecalcualtePayments(cartItemCreated.CartID, c)

	responses.JSON(w, http.StatusCreated, cartItemCreated)
}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(cartID uint, c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Recalcualte the payments

	var cartItems []models.CartItem
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	var totalDiscountedCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
		totalDiscountedCost += item.ItemDiscountedTotal
	}
	var cart models.Cart
	if err := db.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}
	db.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
}
