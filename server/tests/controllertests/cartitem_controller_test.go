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
	seedFruits()
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
		handler := http.HandlerFunc(server.CreateUpdateItemInCart)
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
