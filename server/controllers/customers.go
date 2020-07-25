package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary Creates Customer record
// @Description This end point will record customer details into the database
// @Accept  json
// @Produce  json
// @Param Input body models.CreateCustomerInput true "Input request"
// @Success 200 {object} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers [post]
// CreateCustomer will created customer for the fruit store
func CreateCustomer(c *gin.Context) {
	// Validate input
	db := c.MustGet("db").(*gorm.DB)
	var input models.CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update cart to cart array in the customer table
	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}

	customer := models.Customer{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		LoginId:   input.LoginId,
		Cart:      newcart,
	}

	if err := db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists with the same login id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @Summary Show details of a customer
// @Description Get details of a customer
// @Accept  json
// @Produce  json
// @Param id path string true "Customer identifier"
// @Success 200 {array} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers/{id} [get]
// FindCustomer will return a customer based on the input
func FindCustomer(c *gin.Context) {
	var customer models.Customer
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	var cart models.Cart
	db.Where("customer_id = ? AND status = ?", customer.ID, "OPEN").Find(&cart)
	var cartItem []models.CartItem
	db.Where("cart_id = ?", cart.ID).Find(&cartItem)
	var payment models.Payment
	db.Where("cart_id = ?", cart.ID).Find(&payment)
	var appliedDualItemDiscount []models.AppliedDualItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedDualItemDiscount)
	var appliedSingleItemDiscount []models.AppliedSingleItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemDiscount)
	var appliedSingleItemCoupon []models.AppliedSingleItemCoupon
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemCoupon)
	customer.Cart = cart
	customer.Cart.CartItem = cartItem
	customer.Cart.Payment = payment
	customer.Cart.AppliedDualItemDiscount = appliedDualItemDiscount
	customer.Cart.AppliedSingleItemCoupon = appliedSingleItemCoupon
	customer.Cart.AppliedSingleItemDiscount = appliedSingleItemDiscount

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// @Summary Show details of all customers
// @Description Get details of all customer
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Customer
// @Router /server/api/v1/customers [get]
// FindCustomers will return all customers exists within the fruitshop
func FindCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)

	for _, customer := range customers {
		var cart models.Cart
		db.Where("customer_id = ? AND status = ?", customer.ID, "OPEN").Find(&cart)
		var cartItem []models.CartItem
		db.Where("cart_id = ?", cart.ID).Find(&cartItem)
		var payment models.Payment
		db.Where("cart_id = ?", cart.ID).Find(&payment)
		var appliedDualItemDiscount []models.AppliedDualItemDiscount
		db.Where("cart_id = ?", cart.ID).Find(&appliedDualItemDiscount)
		var appliedSingleItemDiscount []models.AppliedSingleItemDiscount
		db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemDiscount)
		var appliedSingleItemCoupon []models.AppliedSingleItemCoupon
		db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemCoupon)
		customer.Cart = cart
		customer.Cart.CartItem = cartItem
		customer.Cart.Payment = payment
		customer.Cart.AppliedDualItemDiscount = appliedDualItemDiscount
		customer.Cart.AppliedSingleItemCoupon = appliedSingleItemCoupon
		customer.Cart.AppliedSingleItemDiscount = appliedSingleItemDiscount
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

type PayInput struct {
	CustomerID uint    `json:"customerid"`
	CartID     uint    `json:"cartid"`
	Amount     float64 `json:"amount" binding:"required"`
	Status     string  `json:"status"`
}
