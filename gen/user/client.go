// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen fruitshop/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	AddEndpoint  goa.Endpoint
	GetEndpoint  goa.Endpoint
	ShowEndpoint goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(add, get, show goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:  add,
		GetEndpoint:  get,
		ShowEndpoint: show,
	}
}

// Add calls the "add" endpoint of the "user" service.
// Add may return the following errors:
//	- "not_found" (type *NotFound): User not found
//	- error: internal error
func (c *Client) Add(ctx context.Context, p *AddPayload) (res *UserManagement, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UserManagement), nil
}

// Get calls the "get" endpoint of the "user" service.
// Get may return the following errors:
//	- "not_found" (type *NotFound): User not found
//	- error: internal error
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *UserManagement, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UserManagement), nil
}

// Show calls the "show" endpoint of the "user" service.
func (c *Client) Show(ctx context.Context) (res UserManagementCollection, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(UserManagementCollection), nil
}