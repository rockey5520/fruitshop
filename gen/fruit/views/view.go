// Code generated by goa v3.1.3, DO NOT EDIT.
//
// fruit views
//
// Command:
// $ goa gen fruitshop/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// FruitManagement is the viewed result type that is projected based on a view.
type FruitManagement struct {
	// Type to project
	Projected *FruitManagementView
	// View to render
	View string
}

// FruitManagementCollection is the viewed result type that is projected based
// on a view.
type FruitManagementCollection struct {
	// Type to project
	Projected FruitManagementCollectionView
	// View to render
	View string
}

// FruitManagementView is a type that runs validations on a projected type.
type FruitManagementView struct {
	// Name is the unique Name of the Fruit.
	Name *string
	// Cost of the Fruit.
	Cost *float64
}

// FruitManagementCollectionView is a type that runs validations on a projected
// type.
type FruitManagementCollectionView []*FruitManagementView

var (
	// FruitManagementMap is a map of attribute names in result type
	// FruitManagement indexed by view name.
	FruitManagementMap = map[string][]string{
		"default": []string{
			"Name",
			"Cost",
		},
	}
	// FruitManagementCollectionMap is a map of attribute names in result type
	// FruitManagementCollection indexed by view name.
	FruitManagementCollectionMap = map[string][]string{
		"default": []string{
			"Name",
			"Cost",
		},
	}
)

// ValidateFruitManagement runs the validations defined on the viewed result
// type FruitManagement.
func ValidateFruitManagement(result *FruitManagement) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateFruitManagementView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateFruitManagementCollection runs the validations defined on the viewed
// result type FruitManagementCollection.
func ValidateFruitManagementCollection(result FruitManagementCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateFruitManagementCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateFruitManagementView runs the validations defined on
// FruitManagementView using the "default" view.
func ValidateFruitManagementView(result *FruitManagementView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Name", "result"))
	}
	if result.Cost == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Cost", "result"))
	}
	return
}

// ValidateFruitManagementCollectionView runs the validations defined on
// FruitManagementCollectionView using the "default" view.
func ValidateFruitManagementCollectionView(result FruitManagementCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateFruitManagementView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
