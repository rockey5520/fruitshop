package controllers

import "fruitshop/controllers"

func (s *Server) initializeRoutes() {

	// Routes for customer
	s.Router.HandleFunc("/server/api/v1/customers", middlewares.SetMiddlewareJSON(s.FindCustomers)).Methods("GET")
	s.Router.HandleFunc("/server/api/v1/customers/{login_id}", middlewares.SetMiddlewareJSON(s.FindCustomer)).Methods("GET")
	s.Router.HandleFunc("/server/api/v1/customers", middlewares.SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")

	// Routess for fruits
	s.Router.HandleFunc("/server/api/v1/fruits", middlewares.SetMiddlewareJSON(s.FindFruits)).Methods("GET")

	// Routes for cart
	s.Router.HandleFunc("/server/api/v1/cartitem", middlewares.SetMiddlewareJSON(s.CreateUpdateItemInCart)).Methods("POST")
	s.Router.HandleFunc("/server/api/v1/cartitem/{cart_id}", middlewares.SetMiddlewareJSON(s.GetAllCartItems)).Methods("GET")
	s.Router.HandleFunc("/server/api/v1/cart/{cart_id}", middlewares.SetMiddlewareJSON(s.FindCart)).Methods("GET")

	// Route for discounts
	s.Router.HandleFunc("/server/api/v1/discounts/{cart_id}", middlewares.SetMiddlewareJSON(s.FindDiscounts)).Methods("GET")

	// Route for coupon
	s.Router.HandleFunc("/server/api/v1/orangecoupon/{cart_id}/{fruit_id}/", middlewares.SetMiddlewareJSON(s.ApplyTimeSensitiveCoupon)).Methods("GET")

	// Route for payment
	s.Router.HandleFunc("/server/api/v1/pay", middlewares.SetMiddlewareJSON(s.Pay)).Methods("POST")
}
