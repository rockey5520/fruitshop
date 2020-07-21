package models

import "github.com/jinzhu/gorm"

/*Fruit is a meta table of fruits in the inventory
swagger:model Fruit
*/
type Fruit struct {
	// Primary key, created_at, deleted_at, updated_at for each fruit
	gorm.Model
	// Name of the fruit
	Name string `json:"name" gorm:"primary_key;unique_index"`
	// Price of each fruit
	Price float64 `json:"price"`
	// Single Item Discount
	SingleItemDiscount []SingleItemDiscount
	// Single Item Coupon
	SingleItemCoupon []SingleItemCoupon
}
