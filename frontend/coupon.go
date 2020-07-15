package fruitshop

import (
	"context"
	coupon "fruitshop/frontend/gen/coupon"
	"log"
)

// coupon service example implementation.
// The example methods log the requests and return zero values.
type couponsrvc struct {
	logger *log.Logger
}

// NewCoupon returns the coupon service implementation.
func NewCoupon(logger *log.Logger) coupon.Service {
	return &couponsrvc{logger}
}

// Add implements add.
func (s *couponsrvc) Add(ctx context.Context, p *coupon.AddPayload) (res *coupon.CouponManagement, err error) {
	res = &coupon.CouponManagement{}
	s.logger.Print("coupon.add")
	return
}
