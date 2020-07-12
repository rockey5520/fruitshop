// Code generated by goa v3.1.3, DO NOT EDIT.
//
// cart service
//
// Command:
// $ goa gen fruitshop/design

package cart

import (
	"context"
	cartviews "fruitshop/gen/cart/views"
)

// The cart service allows to manage the state of the cart
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (err error)
	// Remove implements remove.
	Remove(context.Context, *RemovePayload) (err error)
	// Get implements get.
	Get(context.Context, *GetPayload) (res CartManagementCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "cart"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"add", "remove", "get"}

// AddPayload is the payload type of the cart service add method.
type AddPayload struct {
	// cartId of the user
	CartID string
	// Name of the fruit
	Name string
	// Number of fruits
	Count int
	// Cost of fruits
	CostPerItem *float64
	// Total cost for the item
	TotalCost *float64
}

// RemovePayload is the payload type of the cart service remove method.
type RemovePayload struct {
	// cartId of the user
	CartID string
	// Name of the fruit
	Name string
	// Number of fruits
	Count int
	// Cost of fruits
	CostPerItem *float64
	// Total cost for the item
	TotalCost *float64
}

// GetPayload is the payload type of the cart service get method.
type GetPayload struct {
	// cartId
	CartID string
}

// CartManagementCollection is the result type of the cart service get method.
type CartManagementCollection []*CartManagement

// A CartManagement type describes users cart state.
type CartManagement struct {
	// cartId is the unique id of the User.
	CartID string
	// Name of the fruit
	Name string
	// Number of fruits
	Count int
	// Cost of Each fruit
	CostPerItem float64
	// Total cost of fruits
	TotalCost float64
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

// NewCartManagementCollection initializes result type CartManagementCollection
// from viewed result type CartManagementCollection.
func NewCartManagementCollection(vres cartviews.CartManagementCollection) CartManagementCollection {
	return newCartManagementCollection(vres.Projected)
}

// NewViewedCartManagementCollection initializes viewed result type
// CartManagementCollection from result type CartManagementCollection using the
// given view.
func NewViewedCartManagementCollection(res CartManagementCollection, view string) cartviews.CartManagementCollection {
	p := newCartManagementCollectionView(res)
	return cartviews.CartManagementCollection{Projected: p, View: "default"}
}

// newCartManagementCollection converts projected type CartManagementCollection
// to service type CartManagementCollection.
func newCartManagementCollection(vres cartviews.CartManagementCollectionView) CartManagementCollection {
	res := make(CartManagementCollection, len(vres))
	for i, n := range vres {
		res[i] = newCartManagement(n)
	}
	return res
}

// newCartManagementCollectionView projects result type
// CartManagementCollection to projected type CartManagementCollectionView
// using the "default" view.
func newCartManagementCollectionView(res CartManagementCollection) cartviews.CartManagementCollectionView {
	vres := make(cartviews.CartManagementCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCartManagementView(n)
	}
	return vres
}

// newCartManagement converts projected type CartManagement to service type
// CartManagement.
func newCartManagement(vres *cartviews.CartManagementView) *CartManagement {
	res := &CartManagement{}
	if vres.CartID != nil {
		res.CartID = *vres.CartID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Count != nil {
		res.Count = *vres.Count
	}
	if vres.CostPerItem != nil {
		res.CostPerItem = *vres.CostPerItem
	}
	if vres.TotalCost != nil {
		res.TotalCost = *vres.TotalCost
	}
	return res
}

// newCartManagementView projects result type CartManagement to projected type
// CartManagementView using the "default" view.
func newCartManagementView(res *CartManagement) *cartviews.CartManagementView {
	vres := &cartviews.CartManagementView{
		CartID:      &res.CartID,
		Name:        &res.Name,
		Count:       &res.Count,
		CostPerItem: &res.CostPerItem,
		TotalCost:   &res.TotalCost,
	}
	return vres
}
