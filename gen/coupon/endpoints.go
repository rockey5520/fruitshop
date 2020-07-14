// Code generated by goa v3.1.3, DO NOT EDIT.
//
// coupon endpoints
//
// Command:
// $ goa gen fruitshop/design

package coupon

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "coupon" service endpoints.
type Endpoints struct {
	Add goa.Endpoint
}

// NewEndpoints wraps the methods of the "coupon" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add: NewAddEndpoint(s),
	}
}

// Use applies the given middleware to all the "coupon" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "coupon".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		res, err := s.Add(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCouponManagement(res, "default")
		return vres, nil
	}
}
