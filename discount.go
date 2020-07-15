package fruitshop

import (
	"context"
	discount "fruitshop/gen/discount"
	"log"
)

// discount service implementation.
// These methods log the requests and return zero values.
type discountsrvc struct {
	logger *log.Logger
}

// NewDiscount returns the discount service implementation.
func NewDiscount(logger *log.Logger) discount.Service {
	return &discountsrvc{logger}
}

// Get implements get feature of the discounts applied for the purchases
func (s *discountsrvc) Get(ctx context.Context, p *discount.GetPayload) (res discount.DiscountManagementCollection, err error) {
	s.logger.Print("discount.show")
	discount := discount.DiscountManagement{
		UserID: p.UserID,
	}
	res, err = getDiscounts(&discount)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.show completed")
	return res, err
}
