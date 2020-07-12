/*
This is the design file. It contains the API specification, methods, inputs, and outputs using Goa DSL code. The objective is to use this as a single source of truth for the entire API source code.
*/
package design

import (
	. "goa.design/goa/v3/dsl"
)

// Main API declaration for User
var _ = API("fruitshop", func() {
	Title("An api for fruit shop online purchases")
	Description("This api manages online fruit shop with CRUD operations")
	Server("fruitshop", func() {
		Host("localhost", func() {
			URI("http://localhost:8080/api/v1")
		})
	})
})

// User Service declaration with two methods and Swagger API specification file
var _ = Service("user", func() {
	Description("The user service allows access to user members")
	Method("add", func() {
		Payload(func() {
			Field(1, "ID", String, "ID")
			Field(2, "UserName", String, "User Name")
			Required("ID", "UserName")
		})
		Result(Empty)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			POST("/api/v1/user/{ID}")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "ID", String, "ID")
			Required("ID")
		})
		Result(UserManagement)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/api/v1/user/{ID}")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Result(CollectionOf(UserManagement))
		HTTP(func() {
			GET("/api/v1/user")
			Response(StatusOK)
		})
	})
	Files("/openapi.json", "./gen/http/openapi.json")
})

// Fruit Service declaration with get method and Swagger API specification file
var _ = Service("fruit", func() {
	Description("The user service allows access to fruits")

	Method("get", func() {
		Payload(func() {
			Field(1, "Name", String, "Name")
			Field(2, "Cost", Float64, "Cost")
			Required("Name", "Cost")
		})
		Result(FruitManagement)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			GET("/api/v1/fruit/{Name}")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Result(CollectionOf(FruitManagement))
		HTTP(func() {
			GET("/api/v1/fruit")
			Response(StatusOK)
		})
	})
	//Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("cart", func() {
	Description("The cart service allows to manage the state of the cart")
	Method("add", func() {
		Payload(func() {
			Field(1, "cartId", String, "cartId of the user")
			Field(2, "Name", String, "Name of the fruit")
			Field(3, "Count", Int, "Number of fruits")
			Required("cartId", "Name", "Count")
		})
		Result(Empty)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			POST("/api/v1/cart/{cartId}")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "cartId", String, "cartId")
			Required("cartId")
		})
		Result(CollectionOf(CartManagement))
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/api/v1/cart/{cartId}")
			Response(StatusOK)
		})
	})

})

// UserManagement is a custom ResultType used to configure views for our custom type
var UserManagement = ResultType("application/vnd.user", func() {
	Description("A UserManagement type describes a User of e-store.")
	Reference(User)
	TypeName("UserManagement")

	Attributes(func() {
		Attribute("ID", String, "ID is the unique id of the User.", func() {
			Example("1")
		})
		Field(2, "UserName")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("UserName")
	})

	Required("ID")
})

// FruitManagement is a custom ResultType used to configure views for our custom type
var FruitManagement = ResultType("application/vnd.fruit", func() {
	Description("A FruitManagement type describes a Fruit of e-store.")
	Reference(Fruit)
	TypeName("FruitManagement")

	Attributes(func() {
		Attribute("Name", String, "Name is the unique Name of the Fruit.", func() {
			Example("Apple")
		})
		Field(2, "Name")
		Field(3, "Cost")

	})

	View("default", func() {
		Attribute("Name")
		Attribute("Cost")
	})

	Required("Name", "Cost")
})

// CartManagement is a custom ResultType used to configure views for our custom type
var CartManagement = ResultType("application/vnd.cart", func() {
	Description("A CartManagement type describes users cart state.")
	Reference(Cart)
	TypeName("CartManagement")

	Attributes(func() {
		Attribute("cartId", String, "cartId is the unique id of the User.", func() {
			Example("1")
		})

		Field(2, "Name")
		Field(3, "Count")
	})

	View("default", func() {
		Attribute("cartId")
		Attribute("Name")
		Attribute("Count")
	})

	Required("cartId")
	Required("Name")
	Required("Count")
})

// User is the custom type for Users in our database
var User = Type("User", func() {
	Description("User describes a customer of store.")
	Attribute("ID", String, "ID is the unique id of the User.", func() {
		Example("1")
	})
	Attribute("UserName", String, "Name of the User", func() {
		Example("Rakesh Mothukuri")
	})
	Required("ID", "UserName")
})

// User is the custom type for Users in our database
var Fruit = Type("Fruit", func() {
	Description("Fruit describes a fruit of store.")
	Attribute("Name", String, "Name is the unique Name of the Fruit.", func() {
		Example("Apple")
	})
	Attribute("Cost", Float64, "Cost of the Fruit.", func() {
		Example(1.0)
	})
	Required("Name", "Cost")
})

// User is the custom type for Users in our database
var Cart = Type("Cart", func() {
	Description("Cart describes a customer cart in the e-store.")
	Attribute("cartId", String, "cartId is the unique id of the User.", func() {
		Example("1")
	})
	Attribute("Name", String, "Name of the fruit", func() {
		Example("Apple")
	})
	Attribute("Count", Int, "Number of fruits", func() {
		Example(2)
	})

	Required("cartId", "Name", "Count")
})

// NotFound is a custom type where we add the queried field in the response
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when " +
		"the requested data that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Example("Item not found")
	})
	Field(2, "id", String, "ID of missing data")
	Required("message", "id")
})
