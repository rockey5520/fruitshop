package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindCustomer returns customer details if customer exists in the store
func FindCustomer(c *gin.Context) {
	var customer models.Customer

	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
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
