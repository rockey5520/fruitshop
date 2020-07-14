package fruitshop

import (
	"context"
	discount "fruitshop/gen/discount"
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

// Show implements show.
func (s *discountsrvc) Show(ctx context.Context, p *discount.ShowPayload) (res discount.DiscountManagementCollection, err error) {
	s.logger.Print("discount.show")
	return
}
