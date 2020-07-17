package fruitshop

import (
	"context"
	cart "fruitshop/gen/cart"
	"fruitshop/gen/user"
	"log"
)

// cart service implementation.
// These methods log the requests and return zero values.
type cartsrvc struct {
	logger *log.Logger
}

// NewCart returns the cart service implementation.
func NewCart(logger *log.Logger) cart.Service {
	return &cartsrvc{logger}
}

// Add implements add method for the cart item.
func (s *cartsrvc) Add(ctx context.Context, p *cart.AddPayload) (err error) {
	s.logger.Print("cart.add")
	newCart := cart.CartManagement{
		UserID: p.UserID,
		Name:   p.Name,
		Count:  p.Count,
	}
	err = AddItemInCart(&newCart)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("cart.add completed")
	return
}

// Remove implements remove functionality for the item in cart
func (s *cartsrvc) Remove(ctx context.Context, p *cart.RemovePayload) (err error) {
	s.logger.Print("cart.remove")
	newCart := cart.CartManagement{
		Name:  p.Name,
		Count: p.Count,
	}
	err = RemoveItemInCart(&newCart)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("cart.remove completed")
	return
}

// Get implements get functionality to fetch cart of the user
func (s *cartsrvc) Get(ctx context.Context, p *cart.GetPayload) (res cart.CartManagementCollection, err error) {
	s.logger.Print("cart.get")
	user := user.UserManagement{
		UserID: p.UserID,
	}
	res, err = ListAllItemsInCartForId(&user)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("cart.get completed")
	return
}
