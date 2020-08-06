package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
)

//GetCart fetched the information about the cart for a given provided cart id
func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cart_id := vars["cart_id"]

	cart := models.Cart{}
	cartFetched, err := cart.FindCartByCartID(server.DB, cart_id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, cartFetched)
}
