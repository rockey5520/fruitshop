/*
This is the design file. It contains the API specification, methods, inputs, and outputs using Goa DSL code. The objective is to use this as a single source of truth for the entire API source code.
*/
package design

import (
	. "goa.design/goa/v3/dsl"
)

// Main API declaration
var _ = API("users", func() {
	Title("An api for users")
	Description("This api manages users with CRUD operations")
	Server("users", func() {
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
			Field(1, "MobieNumber", String, "MobieNumber")
			Field(2, "UserName", String, "User Name")
			Required("MobieNumber", "UserName")
		})
		Result(Empty)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			POST("/api/v1/user/{MobieNumber}")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "MobieNumber", String, "MobieNumber")
			Required("MobieNumber")
		})
		Result(UserManagement)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/api/v1/user/{MobieNumber}")
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

// UserManagement is a custom ResultType used to configure views for our custom type
var UserManagement = ResultType("application/vnd.user", func() {
	Description("A UserManagement type describes a User of e-store.")
	Reference(User)
	TypeName("UserManagement")

	Attributes(func() {
		Attribute("MobieNumber", String, "MobieNumber is the unique id of the User.", func() {
			Example("rockey5520@gmail.com")
		})
		Field(2, "UserName")
	})

	View("default", func() {
		Attribute("MobieNumber")
		Attribute("UserName")
	})

	Required("MobieNumber")
})

// User is the custom type for Users in our database
var User = Type("User", func() {
	Description("User describes a customer of store.")
	Attribute("MobieNumber", String, "MobieNumber is the unique id of the User.", func() {
		Example("1")
	})
	Attribute("UserName", String, "Name of the User", func() {
		Example("Rakesh Mothukuri")
	})
	Required("MobieNumber", "UserName")
})

// NotFound is a custom type where we add the queried field in the response
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when " +
		"the requested data that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Example("User not found")
	})
	Field(2, "id", String, "ID of missing data")
	Required("message", "id")
})
