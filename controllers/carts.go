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
	models.DB.Where("ID = ?", c.Param("cart_id")).Find(&cart)
	fmt.Println(cart)
	c.JSON(http.StatusOK, gin.H{"data": cart})
}

/*
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

// CreatePayment  ill create default payment entry
func CreatePayment(customer models.Customer) {

	var cart models.Cart
	models.DB.Where("customer_id = ?", customer.ID).Find(&cart)

	payment := models.Payment{
		CartId: cart.ID,
		Amount: 0,
		Status: "NOTPAID",
	}
	if err := models.DB.Create(&payment).Error; err != nil {
		panic("Unable to create payment")
	}

}

// @Summary Applied orange 30 coupon code
// @Description This endpoint applied orange 30 percent discount coupon code
// @Accept  json
// @Produce  json
// @Param login_id path string true "Customer identifier"
// @Success 200 {object} models.Coupon
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/orangecoupon/{login_id} [get]
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
	var coupon models.Coupon
	var fruit models.Fruit
	models.DB.Where("name = ?", cartItem.Name).First(&fruit)
	fmt.Println(cartItem.Name)
	fmt.Println(cartItem.Count)
	fmt.Println(discount.Status)
	if cartItem.Name == "Orange" && cartItem.Count > 0 && discount.Status == "NOTAPPLIED" {
		discountCalculated := ((float64(cartItem.Count) * fruit.Price) / 100) * 30
		updatedTotalCost := cartItem.ItemTotal - discountCalculated
		models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", updatedTotalCost)
		RecalcualtePayments(cart)
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "ORANGE30").Update("status", "APPLIED")
		models.DB.Model(&coupon).Where("cart_id = ?", cart.ID, "ORANGE30").Update("status", "APPLIED")
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", float64(cartItem.Count)*fruit.Price)
	models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "ORANGE30").Update("status", "NOTAPPLIED")
	models.DB.Model(&coupon).Where("cart_id = ?", cart.ID, "ORANGE30").Update("status", "NOTAPPLIED")

	RecalcualtePayments(cart)

}
*/
