// Code generated by goa v3.2.0, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	"fmt"
)

// AddUserPath returns the URL path to the user service add HTTP endpoint.
func AddUserPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/user/%v", userID)
}

// GetUserPath returns the URL path to the user service get HTTP endpoint.
func GetUserPath(userID string) string {
	return fmt.Sprintf("/server/api/v1/user/%v", userID)
}

// ShowUserPath returns the URL path to the user service show HTTP endpoint.
func ShowUserPath() string {
	return "/server/api/v1/user"
}
