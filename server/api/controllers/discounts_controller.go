package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
)

//GetAppliedDiscounts will fetch all the discounts applied on a given cart_id of the customer
func (server *Server) GetAppliedDiscounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartid := vars["cart_id"]
	discount := models.Discount{}

	appliedDiscounts := discount.FindAllDiscounts(server.DB, cartid)

	responses.JSON(w, http.StatusOK, appliedDiscounts)

}
