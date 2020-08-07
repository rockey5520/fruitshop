package controllertests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetCartByCartID(t *testing.T) {

	err := refreshCartTable()
	if err != nil {
		log.Fatal(err)
	}
	cart, err := seedOneCart()
	if err != nil {
		log.Fatal(err)
	}
	userSample := []struct {
		statusCode   int
		status       string
		total        float64
		totalsavings float64
	}{
		{
			statusCode:   200,
			status:       "OPEN",
			total:        cart.Total,
			totalsavings: cart.TotalSavings,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/carts", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"cart_id": "1"})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetCart)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		fmt.Println(responseMap)
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, cart.Total, responseMap["total"])
			assert.Equal(t, cart.TotalSavings, responseMap["totalsavings"])
		}
	}
}
