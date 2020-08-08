package models

import "github.com/jinzhu/gorm"

/*
swagger:model AppliedSingleItemDiscount
*/
// AppliedSingleItemDiscount references an applied single item discount in the table.
type AppliedSingleItemDiscount struct {
	// Primary key, created_at, deleted_at, updated_at for each applied single item discount
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// SingleItemDiscountID is the primary key from the SingleItemDiscount table
	SingleItemDiscountID uint
	// SingleItemDiscountName is the name of the discount applied from the SingletemDiscount table
	SingleItemDiscountName string
	// Amount of the savings applied using this discount
	Savings float64 `json:"savings"`
}
