// Code generated by goa v3.2.0, DO NOT EDIT.
//
// coupon service
//
// Command:
// $ goa gen fruitshop/design

package coupon

import (
	"context"
	couponviews "fruitshop/frontend/gen/coupon/views"
)

// The coupon service allows users to apply coupons
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res *CouponManagement, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "coupon"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"add"}

// AddPayload is the payload type of the coupon service add method.
type AddPayload struct {
	// userId
	UserID string
}

// CouponManagement is the result type of the coupon service add method.
type CouponManagement struct {
	// ID is the unique id of the User.
	UserID string
	// ID is the unique id of the Users coupon.
	ID *string
	// Name of the coupon.
	Name *string
	// status of  Users coupon.
	Status *string
	// Users coupon created date time
	CreateTime *string
}

// NotFound is the type returned when the requested data that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing data
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when the requested data that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return "not_found"
}

// NewCouponManagement initializes result type CouponManagement from viewed
// result type CouponManagement.
func NewCouponManagement(vres *couponviews.CouponManagement) *CouponManagement {
	return newCouponManagement(vres.Projected)
}

// NewViewedCouponManagement initializes viewed result type CouponManagement
// from result type CouponManagement using the given view.
func NewViewedCouponManagement(res *CouponManagement, view string) *couponviews.CouponManagement {
	p := newCouponManagementView(res)
	return &couponviews.CouponManagement{Projected: p, View: "default"}
}

// newCouponManagement converts projected type CouponManagement to service type
// CouponManagement.
func newCouponManagement(vres *couponviews.CouponManagementView) *CouponManagement {
	res := &CouponManagement{}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	return res
}

// newCouponManagementView projects result type CouponManagement to projected
// type CouponManagementView using the "default" view.
func newCouponManagementView(res *CouponManagement) *couponviews.CouponManagementView {
	vres := &couponviews.CouponManagementView{
		UserID: &res.UserID,
	}
	return vres
}
