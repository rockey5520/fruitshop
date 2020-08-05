package controllertests

import (
	"encoding/json"
	"fmt"
	"fruitshop/api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetDiscounts(t *testing.T) {

	err := refreshDiscountsTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedSingleItemDiscount()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/discounts", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetDiscounts)
	handler.ServeHTTP(rr, req)

	var discounts []models.Discount
	err = json.Unmarshal([]byte(rr.Body.String()), &discounts)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	fmt.Println(discounts)
	assert.Equal(t, len(discounts), 1)
}
