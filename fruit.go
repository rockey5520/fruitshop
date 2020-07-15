package fruitshop

import (
	"context"
	fruit "fruitshop/gen/fruit"
	"log"
)

// fruit service implementation.
// These methods log the requests and return zero values.
type fruitsrvc struct {
	logger *log.Logger
}

// NewFruit returns the fruit service implementation.
func NewFruit(logger *log.Logger) fruit.Service {

	return &fruitsrvc{logger}
}

// Get implements get feature for the fruits available in the shop per fruit
func (s *fruitsrvc) Get(ctx context.Context, p *fruit.GetPayload) (res *fruit.FruitManagement, err error) {
	res = &fruit.FruitManagement{}
	s.logger.Print("fruit.get")
	return
}

// Show implements show feature for the fruits available in the shop
func (s *fruitsrvc) Show(ctx context.Context) (res fruit.FruitManagementCollection, err error) {
	s.logger.Print("fruit.show started")
	res, err = ListFruits()
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("fruit.show completed")
	return
}
