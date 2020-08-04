package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
)

//GetCart is
func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cartId := vars["cart_id"]

	cart := models.Cart{}
	customerFetched, err := cart.FindCartByCartID(server.DB, cartId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerFetched)
}
