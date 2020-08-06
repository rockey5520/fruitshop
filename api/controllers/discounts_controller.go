package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
)

//GetCustomer is
func (server *Server) GetAppliedDiscounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartid := vars["cart_id"]
	discount := models.Discount{}

	appliedDiscounts := discount.FindAllDiscounts(server.DB, cartid)

	responses.JSON(w, http.StatusOK, appliedDiscounts)

}
