package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"
	"time"

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

//CreateCart will create a cart item for the customer
func CreateCart(customer models.Customer) {
	models.DB.Where("login_id = ?", customer.LoginId).Find(&customer)

	cart := models.Cart{CustomerId: customer.ID, Total: 0.0}

	if err := models.DB.Create(&cart).Error; err != nil {
		panic("Unable to create cart")
	}

}

//CreateCoupon will create a coupon for a customer with default values
func CreateCoupon(customer models.Customer) {
	var cart models.Cart
	models.DB.Where("login_id = ?", customer.LoginId).Find(&customer)
	models.DB.Where("customer_id = ?", customer.ID).Find(&cart)

	coupon := models.Coupon{CartID: cart.ID, Name: "ORANGE30", Status: "NOTAPPLIED"}

	if err := models.DB.Create(&coupon).Error; err != nil {
		panic("Unable to create coupon")
	}

}

//ApplyOrangeCoupon applies discount for oranges
func ApplyOrangeCoupon(c *gin.Context) {
	go ApplyOrange30Coupon(c)
}

func ApplyOrange30Coupon(c *gin.Context) {
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
	var cartItem models.CartItem
	if err := models.DB.Where("cart_id = ? and name = ?", cart.ID, "Orange").Find(&cartItem).Error; err != nil {
		fmt.Println("Error ", err)
	}

	var discount models.Discount
	models.DB.Where("customer_id = ? AND name = ?", cart.CustomerId, "ORANGE30").Find(&discount)
	var fruit models.Fruit
	models.DB.Where("name = ?", cartItem.Name).First(&fruit)
	if cartItem.Name == "Orange" && cartItem.Count > 0 && discount.Status == "NOTAPPLIED" {
		discountCalculated := ((float64(cartItem.Count) * fruit.Price) / 100) * 30
		updatedTotalCost := cartItem.ItemTotal - discountCalculated
		models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", updatedTotalCost)
		RecalcualtePayments(cart)
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "ORANGE30").Update("status", "APPLIED")
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", float64(cartItem.Count)*fruit.Price)
	models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "ORANGE30").Update("status", "NOTAPPLIED")

	RecalcualtePayments(cart)

}
