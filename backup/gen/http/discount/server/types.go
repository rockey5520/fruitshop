// Code generated by goa v3.2.0, DO NOT EDIT.
//
// discount HTTP server types
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	discount "fruitshop/gen/discount"
	discountviews "fruitshop/gen/discount/views"
)

// DiscountManagementResponseCollection is the type of the "discount" service
// "get" endpoint HTTP response body.
type DiscountManagementResponseCollection []*DiscountManagementResponse

// DiscountManagementResponse is used to define fields on response body types.
type DiscountManagementResponse struct {
	// userId for the customer
	UserID string `form:"userId" json:"userId" xml:"userId"`
	// Name of the discount
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Status of the discount
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewDiscountManagementResponseCollection builds the HTTP response body from
// the result of the "get" endpoint of the "discount" service.
func NewDiscountManagementResponseCollection(res discountviews.DiscountManagementCollectionView) DiscountManagementResponseCollection {
	body := make([]*DiscountManagementResponse, len(res))
	for i, val := range res {
		body[i] = marshalDiscountviewsDiscountManagementViewToDiscountManagementResponse(val)
	}
	return body
}

// NewGetPayload builds a discount service get endpoint payload.
func NewGetPayload(userID string) *discount.GetPayload {
	v := &discount.GetPayload{}
	v.UserID = userID

	return v
}