package models

import "github.com/jinzhu/gorm"

type Cart struct {
	// Primary key, created_at, deleted_at, updated_at for each cart
	gorm.Model
	// Foriegn key for the Cart table coming from the Customer table
	CustomerId uint `gorm:"not null"`
	// Total amount valued for the cart
	Total        float64 `json:"total"`
	TotalSavings float64 `json:"totalsavings"`
	// Status of the cart can be either open or closed based on the payment status
	Status string `json:"status"`
	// CartItem is having has-many relationship with Cart
	CartItem []CartItem `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// Payment is having has-one relation with Cart
	Payment Payment `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedDualItemDiscount is having has-many relationship with Cart
	AppliedDualItemDiscount []AppliedDualItemDiscount `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedDualItemDiscount is having has-many relationship with Cart
	AppliedSingleItemDiscount []AppliedSingleItemDiscount `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedSingleItemCoupon is having has-many relationship with Cart
	AppliedSingleItemCoupon []AppliedSingleItemCoupon `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	//ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
}
