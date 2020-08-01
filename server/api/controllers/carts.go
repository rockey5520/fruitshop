package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
func (server *Server) FindCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cart_id, err := vars["cart_id"]
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	cart := models.Cart{}
	cartFetched, err := cart.FindCartByID(server.DB, cart_id)
	responses.JSON(w, http.StatusOK, cartFetched)
}
