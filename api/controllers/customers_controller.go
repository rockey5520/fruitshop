package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"

	"github.com/gorilla/mux"
)

// CreateCustomer is
func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am her ")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	createdCustomer, err := customer.SaveCustomer(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCustomer.ID))
	responses.JSON(w, http.StatusCreated, createdCustomer)
}

//GetCustomer is
func (server *Server) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	loginid := vars["loginid"]

	customer := models.Customer{}
	customerFetched, err := customer.FindCustomerByID(server.DB, loginid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerFetched)
}
