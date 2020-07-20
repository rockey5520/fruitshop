package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Show details of all discounts for a given cart
// @Description Get details of all discounts for each cart
// @Accept  json
// @Produce  json
// @Param login_id path string true "Customer identifier"
// @Success 200 {array} models.Discount
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/discounts/{login_id} [get]

// FindDiscounts will return all discounts with status as APPLIED available within the fruitshop
func FindDiscounts(c *gin.Context) {

	customer := models.Customer{
		LoginId: c.Param("login_id"),
	}
	var discounts []models.Discount

	models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).
		Preload("Discounts").
		Find(&customer)
	discounts = customer.Discounts

	c.JSON(http.StatusOK, gin.H{"data": discounts})
}
