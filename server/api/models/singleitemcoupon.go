package models

import (
	"github.com/jinzhu/gorm"
)

/*
swagger:model SingleItemCoupon
*/
//SingleItemCoupon table carries information of all time sensitive coupons with time it expires in seconds
type SingleItemCoupon struct {
	// Primary key, created_at, deleted_at, updated_at for each SingleItemCoupon
	gorm.Model
	// Name of the Discount
	Name string `json:"name"`
	// Foriegn key for the SingleItemDiscount table coming from the Fruit table
	FruitID uint
	// Percentage of the discount needs to be applied
	Discount int `json:"discount"`
	// duration coupon needs to expire post applying on the cart item (in seconds)
	Duration int
}
