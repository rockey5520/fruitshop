package models

import "github.com/jinzhu/gorm"

/*AppliedSingleItemDiscount is discount asssociated with fruits which are applied for the cart
swagger:model AppliedSingleItemCoupon
*/
type AppliedSingleItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each applied dual item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// SingleItemDiscountID is the primary key from the DualItemDiscount table
	//SingleItemDiscountID uint
	SingleItemDiscount []SingleItemDiscount `gorm:"foreignkey:ID;association_foreignkey:ID"`
	// Percentage of the discount needs to be applied
	Savings float64 `json:"savings"`
}
