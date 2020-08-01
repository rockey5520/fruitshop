package controllers

import (
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary Creates Customer record
// @Description This end point will record customer details into the database
// @Accept  json
// @Produce  json
// @Param Input body models.CreateCustomerInput true "Input request"
// @Success 200 {object} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers [post]
// CreateCustomer will created customer for the fruit store
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

// @Summary Show details of a customer
// @Description Get details of a customer
// @Accept  json
// @Produce  json
// @Param id path string true "Customer identifier"
// @Success 200 {array} models.Customer
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/customers/{id} [get]
// FindCustomer will return a customer based on the input
func (server *Server) FindCustomer(w http.ResponseWriter, r *http.Request) {

	c.JSON(http.StatusOK, gin.H{"data": customer})

	vars := mux.Vars(r)
	login_id := vars["login_id"]

	customer := models.Customer{}
	customerFetched, err := customer.FindCustomerByID(server.DB, login_id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, customerFetched)
}

type PayInput struct {
	CustomerID uint    `json:"customerid"`
	CartID     uint    `json:"cartid"`
	Amount     float64 `json:"amount" binding:"required"`
	Status     string  `json:"status"`
}
