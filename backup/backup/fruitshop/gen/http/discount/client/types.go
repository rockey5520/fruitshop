// Code generated by goa v3.2.0, DO NOT EDIT.
//
// discount HTTP client types
//
// Command:
// $ goa gen fruitshop/design

package client

import (
	discountviews "fruitshop/gen/discount/views"

	goa "goa.design/goa/v3/pkg"
)

// GetResponseBody is the type of the "discount" service "get" endpoint HTTP
// response body.
type GetResponseBody []*DiscountManagementResponse

// DiscountManagementResponse is used to define fields on response body types.
type DiscountManagementResponse struct {
	// userId for the customer
	UserID *string `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
	// Name of the discount
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Status of the discount
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewGetDiscountManagementCollectionOK builds a "discount" service "get"
// endpoint result from a HTTP "OK" response.
func NewGetDiscountManagementCollectionOK(body GetResponseBody) discountviews.DiscountManagementCollectionView {
	v := make([]*discountviews.DiscountManagementView, len(body))
	for i, val := range body {
		v[i] = unmarshalDiscountManagementResponseToDiscountviewsDiscountManagementView(val)
	}
	return v
}

// ValidateDiscountManagementResponse runs the validations defined on
// DiscountManagementResponse
func ValidateDiscountManagementResponse(body *DiscountManagementResponse) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userId", "body"))
	}
	return
}