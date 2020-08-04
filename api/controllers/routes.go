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

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

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
