package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"gopkg.in/go-playground/assert.v1"
)

func TestGetCartByCartID(t *testing.T) {

	err := refreshCartTable()
	if err != nil {
		log.Fatalf("Error refreshCartTable %v\n", err)
	}

	cart, err := seedOneCart()
	if err != nil {
		log.Fatalf("cannot seedOneCustomer: %v", err)
	}
	fountCart, err := cart.FindCartByCartID(server.DB, "1")
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, fountCart.Total, cart.Total)
	assert.Equal(t, fountCart.TotalSavings, cart.TotalSavings)
}
