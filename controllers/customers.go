package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindCustomer returns customer details if customer exists in the store
func FindCustomer(c *gin.Context) {
	var customer models.Customer

	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

//CreateCustomer will created customer for the fruit store
func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
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

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// FindCustomers will retuen all customers exists within the fruitshop
func FindCustomers(c *gin.Context) {

	var customers []models.Customer
	models.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

type CreateCustomerInput struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	LoginId   string `json:"loginid" binding:"required"`
}

//LoadDiscountsInventory will load Discount coupons to the database with default status
func LoadDiscountsInventory(customer models.Customer) {

	apple10 := models.Discount{Name: "APPLE10", Status: "APPLIED", CustomerId: customer.ID}
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
