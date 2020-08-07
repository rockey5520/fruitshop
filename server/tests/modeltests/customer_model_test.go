package modeltests

import (
	"log"
	"testing"

	"fruitshop/api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"gopkg.in/go-playground/assert.v1"
)

func TestSaveCustomer(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatalf("Error refreshCustomerTable %v\n", err)
	}
	err = refreshCartTable()
	if err != nil {
		log.Fatalf("Error refreshCartTable %v\n", err)
	}
	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}
	newCustomer := models.Customer{
		FirstName: "Rakesh",
		LastName:  "Mothukuri",
		LoginID:   "rockey5520",
		Cart:      newcart,
	}

	savedCustomer, err := newCustomer.SaveCustomer(server.DB)
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, newCustomer.FirstName, savedCustomer.FirstName)
	assert.Equal(t, newCustomer.LastName, savedCustomer.LastName)
	assert.Equal(t, newCustomer.LoginID, savedCustomer.LoginID)
}

func TestGetCustomerByLoginID(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatalf("Error refreshCustomerTable %v\n", err)
	}

	customer, err := seedOneCustomer()
	if err != nil {
		log.Fatalf("cannot seedOneCustomer: %v", err)
	}
	foundCustomer, err := customer.FindCustomerByLoginID(server.DB, customer.LoginID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundCustomer.FirstName, customer.FirstName)
	assert.Equal(t, foundCustomer.LastName, customer.LastName)
	assert.Equal(t, foundCustomer.LoginID, customer.LoginID)
}
