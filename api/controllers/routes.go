package controllers

import (
	"fmt"
	"fruitshop/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) initializeRoutes() {

	// Customer routes
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")
	s.Router.HandleFunc("/customers/{loginid}", middlewares.SetMiddlewareJSON(s.GetCustomer)).Methods("GET")

	// Fruit routes
	s.Router.HandleFunc("/fruits", middlewares.SetMiddlewareJSON(s.GetFruits)).Methods("GET")

	// CartItem routes
	s.Router.HandleFunc("/cartitem", middlewares.SetMiddlewareJSON(s.CreateUpdateItemInCart)).Methods("POST")

	// Cart route
	s.Router.HandleFunc("/cart/{cart_id}", middlewares.SetMiddlewareJSON(s.GetCart)).Methods("GET")

	// Discounts routes
	s.Router.HandleFunc("/discounts/{cart_id}", middlewares.SetMiddlewareJSON(s.GetDiscounts)).Methods("GET")

	// Coupon route
	s.Router.HandleFunc("/orangecoupon/{cart_id}/{fruit_id}", middlewares.SetMiddlewareJSON(s.ApplyTimeSensitiveCoupon)).Methods("GET")

	// Pay route
	s.Router.HandleFunc("/pay", middlewares.SetMiddlewareJSON(s.Pay)).Methods("POST")

	fmt.Println("Initialized routes are: ")
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	http.Handle("/", s.Router)
}
