package models

import "github.com/jinzhu/gorm"

/*
swagger:model AppliedDualItemDiscount
*/
// AppliedDualItemDiscount are active discount applied per cart for discount rules which are in combination of fruits
type AppliedDualItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each applied dual item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// DualItemDiscountID is the primary key from the DualItemDiscount table
	DualItemDiscountID uint
	// DualItemDiscountName is the name of the discount( Example : PEARBANANA30 )
	DualItemDiscountName string
	// Amount of the savings applied using this coupond
	Savings float64 `json:"savings"`
}
