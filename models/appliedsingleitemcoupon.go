package models

import "github.com/jinzhu/gorm"

/*AppliedSingleItemDiscount is discount asssociated with fruits which are applied for the cart
swagger:model AppliedSingleItemDiscount
*/
type AppliedSingleItemCoupon struct {
	// Primary key, created_at, deleted_at, updated_at for each applied single item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// SingleItemCouponID is the primary key from the SingleItemCouponID table
	SingleItemCouponID uint
	// Percentage of the discount needs to be applied
	Savings float64 `json:"savings"`
}
