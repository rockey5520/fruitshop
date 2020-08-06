package models

import "github.com/jinzhu/gorm"

/*AppliedDualItemDiscount is discount asssociated with combination of fruits which are applied for the cart
swagger:model AppliedDualItemDiscount
*/
// AppliedDualItemDiscount is
type AppliedDualItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each applied dual item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// DualItemDiscountID is the primary key from the DualItemDiscount table
	//DualItemDiscountID uint
	//DualItemDiscount []DualItemDiscount `gorm:"foreignkey:ID;association_foreignkey:ID"`
	DualItemDiscountID   uint
	DualItemDiscountName string
	// Percentage of the discount needs to be applied
	Savings float64 `json:"savings"`
}
