// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payment HTTP server types
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	payment "fruitshop/gen/payment"
	paymentviews "fruitshop/gen/payment/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "payment" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// cartId of the user
	CartID *string `form:"cartId,omitempty" json:"cartId,omitempty" xml:"cartId,omitempty"`
	// Total cost of the cart
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// AddResponseBody is the type of the "payment" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// cartId is the unique cart id of the User.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// cartId is the unique cart id of the User.
	CartID string `form:"cartId" json:"cartId" xml:"cartId"`
	// Amount to be paid for the purchase
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Payment status
	PaymentStatus *string `form:"paymentStatus,omitempty" json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
}

// GetResponseBody is the type of the "payment" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// cartId is the unique cart id of the User.
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// cartId is the unique cart id of the User.
	CartID string `form:"cartId" json:"cartId" xml:"cartId"`
	// Amount to be paid for the purchase
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Payment status
	PaymentStatus *string `form:"paymentStatus,omitempty" json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "payment" service.
func NewAddResponseBody(res *paymentviews.PaymentManagementView) *AddResponseBody {
	body := &AddResponseBody{
		ID:            res.ID,
		CartID:        *res.CartID,
		Amount:        res.Amount,
		PaymentStatus: res.PaymentStatus,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "payment" service.
func NewGetResponseBody(res *paymentviews.PaymentManagementView) *GetResponseBody {
	body := &GetResponseBody{
		ID:            res.ID,
		CartID:        *res.CartID,
		Amount:        res.Amount,
		PaymentStatus: res.PaymentStatus,
	}
	return body
}

// NewAddPayload builds a payment service add endpoint payload.
func NewAddPayload(body *AddRequestBody, id string) *payment.AddPayload {
	v := &payment.AddPayload{
		CartID: *body.CartID,
		Amount: *body.Amount,
	}
	v.ID = id

	return v
}

// NewGetPayload builds a payment service get endpoint payload.
func NewGetPayload(id string, cartID string) *payment.GetPayload {
	v := &payment.GetPayload{}
	v.ID = id
	v.CartID = cartID

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.CartID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cartId", "body"))
	}
	if body.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("amount", "body"))
	}
	return
}
