package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
)

//GetCart fetches the information about the cart for a given provided cart id
func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {

	// Reading cart_id from request params
	vars := mux.Vars(r)
	cart_id := vars["cart_id"]

	cart := models.Cart{}
	// Fetch cart from DB using provided CartID
	cartFetched, err := cart.FindCartByCartID(server.DB, cart_id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, cartFetched)
}
