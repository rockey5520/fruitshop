// Code generated by goa v3.1.3, DO NOT EDIT.
//
// cart client
//
// Command:
// $ goa gen fruitshop/design

package cart

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "cart" service client.
type Client struct {
	AddEndpoint    goa.Endpoint
	RemoveEndpoint goa.Endpoint
	GetEndpoint    goa.Endpoint
}

// NewClient initializes a "cart" service client given the endpoints.
func NewClient(add, remove, get goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:    add,
		RemoveEndpoint: remove,
		GetEndpoint:    get,
	}
}

// Add calls the "add" endpoint of the "cart" service.
// Add may return the following errors:
//	- "not_found" (type *NotFound): Fruit not found
//	- error: internal error
func (c *Client) Add(ctx context.Context, p *AddPayload) (err error) {
	_, err = c.AddEndpoint(ctx, p)
	return
}

// Remove calls the "remove" endpoint of the "cart" service.
// Remove may return the following errors:
//	- "not_found" (type *NotFound): Fruit not found
//	- error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}

// Get calls the "get" endpoint of the "cart" service.
// Get may return the following errors:
//	- "not_found" (type *NotFound): User not found
//	- error: internal error
func (c *Client) Get(ctx context.Context, p *GetPayload) (res CartManagementCollection, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(CartManagementCollection), nil
}
