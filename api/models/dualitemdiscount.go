package models

import "github.com/jinzhu/gorm"

/*
swagger:model Discount
*/
//DualItemDiscount is a meta table represents the information of available dual item discounts
type DualItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each discount
	gorm.Model
	// Name of the Discount
	Name    string `json:"name"`
	FruitID uint
	// Foriegn key for the DualItemDiscount table coming from the Fruit table
	FruitID_1 uint
	// Foriegn key for the DualItemDiscount table coming from the Fruit table
	FruitID_2 uint
	// Number of items on which discount needs to be applied
	Count_1 int `json:"count1"`
	// Number of items on which discount needs to be applied
	Count_2 int `json:"count2"`
	// Percentage of the discount needs to be applied
	Discount int `json:"discount"`
}
