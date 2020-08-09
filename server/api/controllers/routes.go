package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//initializeRoutes function will initialize http routes for the application using gorilla mux
func (s *Server) initializeRoutes() {

	// Customer routes
	s.Router.HandleFunc("/server/customers", SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")
	s.Router.HandleFunc("/server/customers/{loginid}", SetMiddlewareJSON(s.GetCustomer)).Methods("GET")

	// Fruit routes
	s.Router.HandleFunc("/server/fruits", SetMiddlewareJSON(s.GetFruits)).Methods("GET")

	// CartItem routes
	s.Router.HandleFunc("/server/cartitem", SetMiddlewareJSON(s.CreateItemInCart)).Methods("POST")
	s.Router.HandleFunc("/server/cartitem", SetMiddlewareJSON(s.UpdateItemInCart)).Methods("PUT")
	s.Router.HandleFunc("/server/cartitem/{cart_id}/{fruitname}", SetMiddlewareJSON(s.DeleteItemInCart)).Methods("DELETE")
	s.Router.HandleFunc("/server/cartitems/{cart_id}", SetMiddlewareJSON(s.GetCartItems)).Methods("GET")

	// Cart route
	s.Router.HandleFunc("/server/cart/{cart_id}", SetMiddlewareJSON(s.GetCart)).Methods("GET")

	// Discounts routes
	s.Router.HandleFunc("/server/discounts/{cart_id}", SetMiddlewareJSON(s.GetAppliedDiscounts)).Methods("GET")

	// Coupon route
	s.Router.HandleFunc("/server/orangecoupon/{cart_id}/{fruit_id}", SetMiddlewareJSON(s.ApplyTimeSensitiveCoupon)).Methods("GET")

	// Pay route
	s.Router.HandleFunc("/server/pay", SetMiddlewareJSON(s.Pay)).Methods("POST")

	// Serves angular application on / endpoint
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist/fruitshop-ui"))) // for docker
	//s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist/fruitshop-ui"))) // for local

	fmt.Println()
	fmt.Println("These below are the initialized routes for the application : ")
	fmt.Println()
	err := s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", s.Router)
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
