package modeltests

import (
	"log"
	"testing"

	"fruitshop/api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"gopkg.in/go-playground/assert.v1"
)

func TestSaveCartItem(t *testing.T) {

	err := refreshCartItemTable()
	if err != nil {
		log.Fatalf("Error refreshCartItemTable %v\n", err)
	}
	newCartItem := models.CartItem{
		CartID:              1,
		FruitID:             1,
		Name:                "Apple",
		Quantity:            10,
		ItemTotal:           10,
		ItemDiscountedTotal: 0.0,
	}

	savedCartItem, err := newCartItem.SaveCartItem(server.DB)
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, savedCartItem.FruitID, newCartItem.FruitID)
	assert.Equal(t, savedCartItem.Quantity, newCartItem.Quantity)
	assert.Equal(t, savedCartItem.ItemTotal, newCartItem.ItemTotal)
}

func TestUpdateCartItem(t *testing.T) {

	err := refreshCartItemTable()
	if err != nil {
		log.Fatalf("Error refreshCartItemTable %v\n", err)
	}

	refreshFruitTable()

	err = seedFruits()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedOneCart()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedOneCartItem()
	if err != nil {
		log.Fatal(err)
	}
	newCartItem := models.CartItem{
		CartID:              1,
		FruitID:             1,
		Name:                "Apple",
		Quantity:            1,
		ItemTotal:           1,
		ItemDiscountedTotal: 0.0,
	}

	savedCartItem, err := newCartItem.UpdateCartItem(server.DB)
	if err != nil {
		t.Errorf("Error while saving a cart: %v\n", err)
		return
	}
	assert.Equal(t, savedCartItem.FruitID, newCartItem.FruitID)
	assert.Equal(t, savedCartItem.Quantity, newCartItem.Quantity)
	assert.Equal(t, savedCartItem.ItemTotal, newCartItem.ItemTotal)
}
