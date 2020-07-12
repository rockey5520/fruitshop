package fruitshop

import (
	"context"
	cart "fruitshop/gen/cart"

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
	newCart := cart.CartManagement{
		CartID: p.CartID,
		Name:   p.Name,
		Count:  p.Count,
	}
	err = UpdateItemInCart(&newCart)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("cart.add completed")
	return
}

// Get implements fetching all items from cart.
func (s *cartsrvc) Get(ctx context.Context, p *cart.GetPayload) (res cart.CartManagementCollection, err error) {
	s.logger.Print("cart.get")
	res, err = ListAllItemsInCartForId(p.CartID)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("cart.get completed")
	return
}
