package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Show details of a cart
// @Description Get details of a cart
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"
// @Success 200 {object} models.Cart
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/cart/{cart_id} [get]
// FindCart will fetch the details about the cart of the customer
func FindCart(c *gin.Context) {
	cart := models.Cart{}
	//models.DB.Where("ID = ?", c.Param("cart_id")).Find(&cart)
	models.DB.Where("ID = ?", c.Param("cart_id")).
		Preload("CartItem").
		Preload("Payment").
		Preload("AppliedDualItemDiscount").
		Preload("AppliedSingleItemDiscount").
		Preload("AppliedSingleItemCoupon").
		Find(&cart)

	fmt.Println(cart)
	c.JSON(http.StatusOK, gin.H{"data": cart})
}
