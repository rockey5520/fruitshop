package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// CreateCustomer will created customer for the fruit store
// @Summary Creates Customer record
// @Description This end point will record customer details into the database
// @Accept  json
// @Produce  json
// @Param Input body models.CreateCustomerInput true "Input request"
// @Success 200 {object} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers [post]
func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//user.Prepare()
	// Validate input
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusCreated, customerCreated)
}

// FindCustomer will return a customer based on the input
// @Summary Show details of a customer
// @Description Get details of a customer
// @Accept  json
// @Produce  json
// @Param id path string true "Customer identifier"
// @Success 200 {array} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers/{id} [get]
func (server *Server) FindCustomer(w http.ResponseWriter, r *http.Request) {

	c.JSON(http.StatusOK, gin.H{"data": customer})

	vars := mux.Vars(r)
	loginID := vars["login_id"]

	customer := models.Customer{}
	customerFetched, err := customer.FindCustomerByID(server.DB, login_id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerFetched)
}

//PayInput is
type PayInput struct {
	CustomerID uint    `json:"customerid"`
	CartID     uint    `json:"cartid"`
	Amount     float64 `json:"amount" binding:"required"`
	Status     string  `json:"status"`
}
