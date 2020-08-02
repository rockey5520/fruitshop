package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"
)

// CreateCustomer is
func (server *Server) CreateUpdateItemInCart(w http.ResponseWriter, r *http.Request) {
	// Reading the request body from http request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// Creating customer , cart structs and mapping request body to customer and a creating new card with customer ID
	cartItem := models.CartItem{}
	err = json.Unmarshal(body, &cartItem)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Customer validation
	err = cartItem.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdCustomer, err := cartItem.SaveOrUpdateCartItem(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCustomer.ID))
	responses.JSON(w, http.StatusCreated, createdCustomer)
}
