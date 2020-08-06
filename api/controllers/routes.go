package controllers

import (
	"fmt"
	"fruitshop/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//initializeRoutes function will initialize http routes for the application using gorilla mux
func (s *Server) initializeRoutes() {

	// Customer routes
	s.Router.HandleFunc("/server/customers", middlewares.SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")
	s.Router.HandleFunc("/server/customers/{loginid}", middlewares.SetMiddlewareJSON(s.GetCustomer)).Methods("GET")

	// Fruit routes
	s.Router.HandleFunc("/server/fruits", middlewares.SetMiddlewareJSON(s.GetFruits)).Methods("GET")

	// CartItem routes
	s.Router.HandleFunc("/server/cartitem", middlewares.SetMiddlewareJSON(s.CreateUpdateItemInCart)).Methods("POST")
	s.Router.HandleFunc("/server/cartitems/{cart_id}", middlewares.SetMiddlewareJSON(s.GetCartItems)).Methods("GET")

	// Cart route
	s.Router.HandleFunc("/server/cart/{cart_id}", middlewares.SetMiddlewareJSON(s.GetCart)).Methods("GET")

	// Discounts routes
	s.Router.HandleFunc("/server/discounts/{cart_id}", middlewares.SetMiddlewareJSON(s.GetAppliedDiscounts)).Methods("GET")

	// Coupon route
	s.Router.HandleFunc("/server/orangecoupon/{cart_id}/{fruit_id}", middlewares.SetMiddlewareJSON(s.ApplyTimeSensitiveCoupon)).Methods("GET")

	// Pay route
	s.Router.HandleFunc("/server/pay", middlewares.SetMiddlewareJSON(s.Pay)).Methods("POST")

	// Serves angular application on / endpoint
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist/fruitshop-ui")))

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
