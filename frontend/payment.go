package fruitshop

import (
	"context"
	payment "fruitshop/frontend/gen/payment"
	"log"
)

// payment service example implementation.
// The example methods log the requests and return zero values.
type paymentsrvc struct {
	logger *log.Logger
}

// NewPayment returns the payment service implementation.
func NewPayment(logger *log.Logger) payment.Service {
	return &paymentsrvc{logger}
}

// Add implements add.
func (s *paymentsrvc) Add(ctx context.Context, p *payment.AddPayload) (res *payment.PaymentManagement, err error) {
	res = &payment.PaymentManagement{}
	s.logger.Print("payment.add")
	return
}

// Get implements get.
func (s *paymentsrvc) Get(ctx context.Context, p *payment.GetPayload) (res *payment.PaymentManagement, err error) {
	res = &payment.PaymentManagement{}
	s.logger.Print("payment.get")
	return
}
