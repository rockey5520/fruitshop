package models

import "github.com/jinzhu/gorm"

/*AppliedSingleItemDiscount is discount asssociated with fruits which are applied for the cart
swagger:model AppliedSingleItemDiscount
*/
// AppliedSingleItemDiscount references an applied single item discount in the table.
type AppliedSingleItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each applied dual item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// SingleItemDiscountID is the primary key from the DualItemDiscount table
	SingleItemDiscountID uint
	// SingleItemDiscountName is the name of the discount applied from the DualItemDiscount table
	SingleItemDiscountName string
	// Percentage of the discount needs to be applied
	Savings float64 `json:"savings"`
}
