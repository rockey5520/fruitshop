package models

import (
	"github.com/jinzhu/gorm"
)

/*SingleItemDiscount is asssociated with Fruit with Has-many relationship
swagger:model Discount
*/
//SingleItemCoupon is
type SingleItemCoupon struct {
	// Primary key, created_at, deleted_at, updated_at for each SingleItemCoupon
	gorm.Model
	// Name of the Discount
	Name string `json:"name"`
	// Foriegn key for the SingleItemDiscount table coming from the Fruit table
	FruitID uint
	// Percentage of the discount needs to be applied
	Discount int `json:"discount"`
}
