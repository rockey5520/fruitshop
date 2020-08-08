package models

import (
	"github.com/jinzhu/gorm"
)

/*
swagger:model Discount
*/
// Discount struct represents the applied discounts in the DB for an given cart
type Discount struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the Discount table coming from the Customer table
	CustomerId uint `gorm:"not null"`
	// Name of the coupon
	Name string `json:"name"`
	// Status of the coupon APPLIED and NOTAPPLIED are the two possible states
	Status string `json:"status"`
}

// FindAllDiscounts fetch all applied discounts for a cart based on the cartID ( this is simple fetch and create array and return type function and no computes made)
func (u *Discount) FindAllDiscounts(db *gorm.DB, cartID string) *[]Discount {

	appliedDiscountsResponseList := make([]Discount, 0)
	appliedSingleItemDiscount := []AppliedSingleItemDiscount{}
	db.Where("cart_id = ?", cartID).
		Find(&appliedSingleItemDiscount)

	appliedDualItemDiscount := []AppliedDualItemDiscount{}
	db.Where("cart_id = ?", cartID).
		Find(&appliedDualItemDiscount)

	appliedSingleItemCoupon := []AppliedSingleItemCoupon{}
	db.Where("cart_id = ?", cartID).
		Find(&appliedSingleItemCoupon)

	for _, singleItemDiscount := range appliedSingleItemDiscount {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
			Name:   singleItemDiscount.SingleItemDiscountName,
			Status: "APPLIED",
		})
	}
	for _, dualItemDiscount := range appliedDualItemDiscount {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
			Name:   dualItemDiscount.DualItemDiscountName,
			Status: "APPLIED",
		})
	}

	for _, singeItemCoupon := range appliedSingleItemCoupon {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
			Name:   singeItemCoupon.SingleItemCouponName,
			Status: "APPLIED",
		})
	}

	return &appliedDiscountsResponseList
}
