package models

import "github.com/jinzhu/gorm"

/*
swagger:model Discount
*/
//SingleItemDiscount is asssociated with Fruit with Has-many relationship
type SingleItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each single item discount
	gorm.Model
	// Name of the Discount
	Name string `json:"name"`
	// Foriegn key for the SingleItemDiscount table coming from the Fruit table
	FruitID uint
	// Number of items on which discount needs to be applied
	Count int `json:"count"`
	// Percentage of the discount needs to be applied
	Discount int `json:"discount"`
}
