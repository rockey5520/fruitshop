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
			Field(2, "userName", String, "User Name")
			Required("ID", "userName")
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
			Field(1, "name", String, "name")
			Field(2, "cost", Float64, "cost")
			Required("name", "cost")
		})
		Result(FruitManagement)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			GET("/api/v1/fruit/{name}")
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
			Field(2, "name", String, "name of the fruit")
			Field(3, "count", Int, "Number of fruits")
			Field(4, "costPerItem", Float64, "Cost of fruits")
			Field(5, "totalCost", Float64, "Total cost for the item")
			Required("cartId", "name", "count")
		})
		Result(Empty)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			POST("/api/v1/cart/add/{cartId}")
			Response(StatusAccepted)
		})
	})

	Method("remove", func() {
		Payload(func() {
			Field(1, "cartId", String, "cartId of the user")
			Field(2, "name", String, "Name of the fruit")
			Field(3, "count", Int, "Number of fruits")
			Field(4, "costPerItem", Float64, "Cost of fruits")
			Field(5, "totalCost", Float64, "Total cost for the item")
			Required("cartId", "name", "count")
		})
		Result(Empty)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			POST("/api/v1/cart/remove/{cartId}")
			Response(StatusAccepted)
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

var _ = Service("payment", func() {
	Description("The cart service allows to manage the state of the cart")
	Method("add", func() {
		Payload(func() {
			Field(1, "ID", String, "ID of the user")
			Field(2, "cartId", String, "cartId of the user")
			Field(2, "amount", Float64, "Total cost of the cart")
			Required("ID", "cartId", "amount")
		})
		Result(PaymentManagement)
		Error("not_found", NotFound, "Fruit not found")
		HTTP(func() {
			POST("/api/v1/payment/pay/{ID}")
			Response(StatusAccepted)
		})
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "ID", String, "cartId")
			Field(1, "cartId", String, "cartId")
			Required("ID", "cartId")
		})
		Result(PaymentManagement)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/api/v1/payment/{ID}/{cartId}")
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
		Field(2, "userName")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("userName")
	})

	Required("ID")
})

// FruitManagement is a custom ResultType used to configure views for our custom type
var FruitManagement = ResultType("application/vnd.fruit", func() {
	Description("A FruitManagement type describes a Fruit of e-store.")
	Reference(Fruit)
	TypeName("FruitManagement")

	Attributes(func() {
		Attribute("name", String, "Name is the unique Name of the Fruit.", func() {
			Example("Apple")
		})
		Field(2, "name")
		Field(3, "cost")

	})

	View("default", func() {
		Attribute("name")
		Attribute("cost")
	})

	Required("name", "cost")
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

		Field(2, "name")
		Field(3, "count")
		Field(4, "costPerItem")
		Field(5, "totalCost")
	})

	View("default", func() {
		Attribute("cartId")
		Attribute("name")
		Attribute("count")
		Attribute("costPerItem")
		Attribute("totalCost")
	})

	Required("cartId")
	Required("name")
	Required("count")
	Required("costPerItem")
	Required("totalCost")
})

// UserManagement is a custom ResultType used to configure views for our custom type
var PaymentManagement = ResultType("application/vnd.payment", func() {
	Description("A PaymentManagement type for the payment for the fruits purchased")
	Reference(Payment)
	TypeName("PaymentManagement")

	Attributes(func() {
		Attribute("ID", String, "cartId is the unique cart id of the User.", func() {
			Example("1")
		})
		Attribute("cartId", String, "cartId is the unique cart id of the User.", func() {
			Example("1")
		})
		Field(2, "amount")
		Field(3, "paymentStatus")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("cartId")
		Attribute("amount")
		Attribute("paymentStatus")
	})

	Required("cartId")
})

// User is the custom type for Users in our database
var User = Type("User", func() {
	Description("User describes a customer of store.")
	Attribute("ID", String, "ID is the unique id of the User.", func() {
		Example("1")
	})
	Attribute("UuerName", String, "Name of the User", func() {
		Example("Rakesh Mothukuri")
	})
	Required("ID", "userName")
})

// User is the custom type for Users in our database
var Fruit = Type("Fruit", func() {
	Description("Fruit describes a fruit of store.")
	Attribute("name", String, "Name is the unique Name of the Fruit.", func() {
		Example("Apple")
	})
	Attribute("cost", Float64, "Cost of the Fruit.", func() {
		Example(1.0)
	})
	Required("name", "cost")
})

// User is the custom type for Users in our database
var Cart = Type("Cart", func() {
	Description("Cart describes a customer cart in the e-store.")
	Attribute("cartId", String, "cartId is the unique id of the User.", func() {
		Example("1")
	})
	Attribute("name", String, "Name of the fruit", func() {
		Example("Apple")
	})
	Attribute("count", Int, "Number of fruits", func() {
		Example(2)
	})
	Attribute("costPerItem", Float64, "Cost of Each fruit", func() {
		Example(2)
	})
	Attribute("totalCost", Float64, "Total cost of fruits", func() {
		Example(4)
	})

	Required("cartId", "name", "count")
})

// Payment is the custom type for Payment in our database
var Payment = Type("Payment", func() {
	Description("Payment describes payment for the items purchased")
	Attribute("cartId", String, "cartId is the unique cart id of the User.", func() {
		Example("1")
	})
	Attribute("ID", String, "Payment ID for the cart", func() {
		Example("30")
	})
	Attribute("amount", Float64, "Amount to be paid for the purchase", func() {
		Example(50)
	})
	Attribute("paymentStatus", String, "Payment status", func() {
		Example("Success")
	})
	Required("cartId")
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
