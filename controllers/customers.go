package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Show details of a customer
// @Description Get details of a customer
// @Accept  json
// @Produce  json
// @Param login_id path string true "Customer identifier"
// @Success 200 {array} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers/{login_id} [get]
// FindCustomer will return a customer based on the input
func FindCustomer(c *gin.Context) {
	var customer models.Customer

	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

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
	var input models.CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new customer
	customer := models.Customer{FirstName: input.FirstName, LastName: input.LastName, LoginId: input.LoginId}
	if err := models.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists with the same login id"})
		return
	}

	// Loads default discounts for the customer  with default values
	LoadDiscountsInventory(customer)
	//Creates an entry for cart for the customer with default values
	CreateCart(customer)
	// Creates coupon entry for the cart and set with default values
	CreateCoupon(customer)
	// Creates payment entry for the cart and set with default values
	CreatePayment(customer)

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

	var customers []models.Customer
	models.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

type PayInput struct {
	Amount float64 `json:"amount" binding:"required"`
	Status string  `json:"status"`
}

type CartItemInput struct {
	Name  string `json:"name" binding:"required"`
	Count int    `json:"count"`
}

//LoadDiscountsInventory will load Discount coupons to the database with default status
func LoadDiscountsInventory(customer models.Customer) {

	apple10 := models.Discount{Name: "APPLE10", Status: "NOTAPPLIED", CustomerId: customer.ID}
	orange30 := models.Discount{Name: "ORANGE30", Status: "NOTAPPLIED", CustomerId: customer.ID}
	pearbanana30 := models.Discount{Name: "PEARBANANA", Status: "NOTAPPLIED", CustomerId: customer.ID}

	if err := models.DB.Create(&apple10).Error; err != nil {
		panic("Unable to create inventory")
	}
	if err := models.DB.Create(&orange30).Error; err != nil {
		panic("Unable to create inventory")
	}
	if err := models.DB.Create(&pearbanana30).Error; err != nil {
		panic("Unable to create inventory")
	}

}
