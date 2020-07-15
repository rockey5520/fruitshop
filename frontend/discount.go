package fruitshop

import (
	"context"
	discount "fruitshop/frontend/gen/discount"
	"log"
)

// discount service example implementation.
// The example methods log the requests and return zero values.
type discountsrvc struct {
	logger *log.Logger
}

// NewDiscount returns the discount service implementation.
func NewDiscount(logger *log.Logger) discount.Service {
	return &discountsrvc{logger}
}

// Get implements get.
func (s *discountsrvc) Get(ctx context.Context, p *discount.GetPayload) (res discount.DiscountManagementCollection, err error) {
	s.logger.Print("discount.get")
	return
}
