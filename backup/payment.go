package fruitshop

import (
	"context"
	payment "fruitshop/gen/payment"
	"log"
)

// payment service implementation.
// These methods log the requests and return zero values.
type paymentsrvc struct {
	logger *log.Logger
}

// NewPayment returns the payment service implementation.
func NewPayment(logger *log.Logger) payment.Service {
	return &paymentsrvc{logger}
}

// Add implements add feature for payment processing
func (s *paymentsrvc) Add(ctx context.Context, p *payment.AddPayload) (res *payment.PaymentManagement, err error) {
	res = &payment.PaymentManagement{}
	s.logger.Print("payment.add started")
	payment := payment.PaymentManagement{
		UserID: p.UserID,
		ID:     p.ID,
	}

	result, err := PayAmount(&payment)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("payment.get completed")
	return &result, err
}

// Get implements get feature for payment amount fetch for the user
func (s *paymentsrvc) Get(ctx context.Context, p *payment.GetPayload) (res *payment.PaymentManagement, err error) {

	s.logger.Print("payment.get started")
	payment := payment.PaymentManagement{
		UserID: p.UserID,
		ID:     p.ID,
	}

	result, err := GetPaymentAmoutFromCart(&payment)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("payment.get completed")
	return &result, err
}
