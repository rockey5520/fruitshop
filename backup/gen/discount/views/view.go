// Code generated by goa v3.2.0, DO NOT EDIT.
//
// discount views
//
// Command:
// $ goa gen fruitshop/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// DiscountManagementCollection is the viewed result type that is projected
// based on a view.
type DiscountManagementCollection struct {
	// Type to project
	Projected DiscountManagementCollectionView
	// View to render
	View string
}

// DiscountManagementCollectionView is a type that runs validations on a
// projected type.
type DiscountManagementCollectionView []*DiscountManagementView

// DiscountManagementView is a type that runs validations on a projected type.
type DiscountManagementView struct {
	// userId for the customer
	UserID *string
	// Name of the discount
	Name *string
	// Status of the discount
	Status *string
}

var (
	// DiscountManagementCollectionMap is a map of attribute names in result type
	// DiscountManagementCollection indexed by view name.
	DiscountManagementCollectionMap = map[string][]string{
		"default": []string{
			"userId",
			"name",
			"status",
		},
	}
	// DiscountManagementMap is a map of attribute names in result type
	// DiscountManagement indexed by view name.
	DiscountManagementMap = map[string][]string{
		"default": []string{
			"userId",
			"name",
			"status",
		},
	}
)

// ValidateDiscountManagementCollection runs the validations defined on the
// viewed result type DiscountManagementCollection.
func ValidateDiscountManagementCollection(result DiscountManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateDiscountManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateDiscountManagementCollectionView runs the validations defined on
// DiscountManagementCollectionView using the "default" view.
func ValidateDiscountManagementCollectionView(result DiscountManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateDiscountManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDiscountManagementView runs the validations defined on
// DiscountManagementView using the "default" view.
func ValidateDiscountManagementView(result *DiscountManagementView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userId", "result"))
	}
	return
}
