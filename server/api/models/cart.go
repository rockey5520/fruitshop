package models

import (
	"github.com/jinzhu/gorm"
)

/*
swagger:model Cart
*/
// Cart represents full information of the cart , items added to the cart, Discounts applied and Payment to be made as a single entity.
type Cart struct {
	// Primary key, created_at, deleted_at, updated_at for each cart
	gorm.Model
	// Foriegn key for the Cart table coming from the Customer table
	CustomerId uint `gorm:"not null"`
	// Total amount valued for the cart
	Total float64 `json:"total"`
	// TotalSavings for the cart
	TotalSavings float64 `json:"totalsavings"`
	// Status of the cart can be either OPEN or CLOSED based on the payment status
	Status string `json:"status"`
	// CartItem is having has-many relationship with Cart
	CartItem []CartItem `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// Payment is having has-one relation with Cart
	Payment Payment `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedDualItemDiscount represents all discounts applied to a cart based on the rule from DualItemDiscount table and is having has-many relationship with Cart
	AppliedDualItemDiscount []AppliedDualItemDiscount `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedSingleItemDiscount represents all discounts applied to a cart based on the rule from SingleItemDiscount table and is having has-many relationship with Cart
	AppliedSingleItemDiscount []AppliedSingleItemDiscount `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// AppliedSingleItemCoupon represents all discounts applied to a cart based on the rule from SingleItemDiscount table  is having has-many relationship with Cart
	AppliedSingleItemCoupon []AppliedSingleItemCoupon `gorm:"foreignkey:CartID;association_foreignkey:ID"`
}

//FindCustomerByID fetches information of a cart by cartID
func (c *Cart) FindCartByCartID(db *gorm.DB, cartID string) (*Cart, error) {
	// Preloading all the tables using cartID which joins all table, This reduces multiple lines of code and
	// represents simple and elegant for anyone reading code for first time
	err := db.Where("ID = ?", cartID).
		Preload("CartItem").
		Preload("Payment").
		Preload("AppliedDualItemDiscount").
		Preload("AppliedSingleItemDiscount").
		Preload("AppliedSingleItemCoupon").
		Find(&c).Error

	return c, err
}
