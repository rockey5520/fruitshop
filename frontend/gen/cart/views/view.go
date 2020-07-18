// Code generated by goa v3.2.0, DO NOT EDIT.
//
// cart views
//
// Command:
// $ goa gen fruitshop/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// CartManagementCollection is the viewed result type that is projected based
// on a view.
type CartManagementCollection struct {
	// Type to project
	Projected CartManagementCollectionView
	// View to render
	View string
}

// CartManagementCollectionView is a type that runs validations on a projected
// type.
type CartManagementCollectionView []*CartManagementView

// CartManagementView is a type that runs validations on a projected type.
type CartManagementView struct {
	// userId is the unique id of the Cart.
	UserID *string
	// Name of the fruit
	Name *string
	// Number of fruits
	Count *int
	// Cost of Each fruit
	CostPerItem *float64
	// Total cost of fruits
	TotalCost *float64
}

var (
	// CartManagementCollectionMap is a map of attribute names in result type
	// CartManagementCollection indexed by view name.
	CartManagementCollectionMap = map[string][]string{
		"default": []string{
			"userId",
			"name",
			"count",
			"costPerItem",
			"totalCost",
		},
	}
	// CartManagementMap is a map of attribute names in result type CartManagement
	// indexed by view name.
	CartManagementMap = map[string][]string{
		"default": []string{
			"userId",
			"name",
			"count",
			"costPerItem",
			"totalCost",
		},
	}
)

// ValidateCartManagementCollection runs the validations defined on the viewed
// result type CartManagementCollection.
func ValidateCartManagementCollection(result CartManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCartManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateCartManagementCollectionView runs the validations defined on
// CartManagementCollectionView using the "default" view.
func ValidateCartManagementCollectionView(result CartManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateCartManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCartManagementView runs the validations defined on
// CartManagementView using the "default" view.
func ValidateCartManagementView(result *CartManagementView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userId", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("count", "result"))
	}
	if result.CostPerItem == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("costPerItem", "result"))
	}
	if result.TotalCost == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("totalCost", "result"))
	}
	return
}