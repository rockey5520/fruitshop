package fruitshop

import (
	"context"
	cart "fruitshop/frontend/gen/cart"
	"log"
)

// cart service example implementation.
// The example methods log the requests and return zero values.
type cartsrvc struct {
	logger *log.Logger
}

// NewCart returns the cart service implementation.
func NewCart(logger *log.Logger) cart.Service {
	return &cartsrvc{logger}
}

// Add implements add.
func (s *cartsrvc) Add(ctx context.Context, p *cart.AddPayload) (err error) {
	s.logger.Print("cart.add")
	return
}

// Remove implements remove.
func (s *cartsrvc) Remove(ctx context.Context, p *cart.RemovePayload) (err error) {
	s.logger.Print("cart.remove")
	return
}

// Get implements get.
func (s *cartsrvc) Get(ctx context.Context, p *cart.GetPayload) (res cart.CartManagementCollection, err error) {
	s.logger.Print("cart.get")
	return
}
