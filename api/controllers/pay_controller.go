package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"
)

//Pay will enable the payment of the money for the given cart
func (server *Server) Pay(w http.ResponseWriter, r *http.Request) {

	// Reading the request body from http request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	payment := models.Payment{}
	err = json.Unmarshal(body, &payment)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer, err := payment.Pay(server.DB, payment)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusCreated, customer)
}
