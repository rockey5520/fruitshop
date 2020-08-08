package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateCustomer(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}
	err = refreshCartTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		firstname    string
		lastname     string
		loginid      string
		errorMessage string
	}{
		{
			inputJSON:    `{"firstname":"Rakesh", "lastname": "mothukuri", "loginid": "a"}`,
			statusCode:   201,
			firstname:    "Rakesh",
			lastname:     "mothukuri",
			loginid:      "a",
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/customers", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateCustomer)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {

			assert.Equal(t, responseMap["firstname"], v.firstname)
			assert.Equal(t, responseMap["lastname"], v.lastname)
		}
	}
}

func TestGetCustomerByLoginID(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}
	err = refreshCartTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneCustomer()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		firstname    string
		lastname     string
		loginid      string
		errorMessage string
	}{
		{
			inputJSON:    `{"firstname":"Rakesh", "lastname": "mothukuri", "loginid": "a"}`,
			statusCode:   200,
			firstname:    "Rakesh",
			lastname:     "mothukuri",
			loginid:      "a",
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("GET", "/customers", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"loginid": v.loginid})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetCustomer)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 201 {
			assert.Equal(t, user.FirstName, responseMap["firstname"])
			assert.Equal(t, user.LastName, responseMap["lastname"])
		}
	}
}

func TestGetCustomerByLoginIDNotAvailable(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}
	err = refreshCartTable()
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		firstname    string
		lastname     string
		loginid      string
		errorMessage string
	}{
		{
			inputJSON:    `{"firstname":"Rakesh", "lastname": "mothukuri", "loginid": "a"}`,
			statusCode:   400,
			firstname:    "Rakesh",
			lastname:     "mothukuri",
			loginid:      "a",
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("GET", "/customers", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"loginid": v.loginid})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetCustomer)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

	}
}
