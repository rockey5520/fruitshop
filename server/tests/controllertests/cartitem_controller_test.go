package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateCartItem(t *testing.T) {

	err := refreshCartItemTable()

	if err != nil {
		log.Fatal(err)
	}
	_, err = seedFruits()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		itemtotal    float64
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 2
			}`,
			statusCode:   201,
			CartID:       1,
			FruitID:      1,
			itemtotal:    2.0,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateItemInCart)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["itemtotal"], v.itemtotal)

		}
	}
}

func TestUpdateCartItem(t *testing.T) {

	err := refreshCartItemTable()

	if err != nil {
		log.Fatal(err)
	}
	_, err = seedFruits()
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
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		quantity     uint
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 2
			}`,
			statusCode:   500,
			CartID:       1,
			FruitID:      1,
			quantity:     4,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("PUT", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateItemInCart)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.NotEqual(t, responseMap["quantity"], v.quantity)

		}
	}
}

func TestDeleteCartItem(t *testing.T) {

	err := refreshCartItemTable()

	if err != nil {
		log.Fatal(err)
	}
	_, err = seedFruits()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedOneCartItem()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedOneCart()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		CartID       uint
		FruitID      uint
		quantity     uint
		errorMessage string
	}{
		{
			inputJSON: `{
				"cartid": 1,
				"name": "Apple",
				"quantity": 0
			}`,
			statusCode:   201,
			CartID:       1,
			FruitID:      1,
			quantity:     4,
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("DELETE", "/cartitem", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteItemInCart)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.NotEqual(t, responseMap["quantity"], v.quantity)

		}
	}
}
