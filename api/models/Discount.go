package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

/*Discount is asssociated with Customer with Has-many relationship
swagger:model Discount
*/
// Discount is
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

// FindAllDiscounts isFindAllDiscounts
func (u *Discount) FindAllDiscounts(db *gorm.DB, cartID string) *[]Discount {

	appliedDiscountsResponseList := make([]Discount, 0)
	appliedSingleItemDiscount := AppliedSingleItemDiscount{}
	db.Where("cart_id = ?", cartID).
		Preload("SingleItemDiscount").
		Find(&appliedSingleItemDiscount)
	fmt.Println(appliedSingleItemDiscount.Savings)
	appliedDualItemDiscount := AppliedDualItemDiscount{}
	db.Where("cart_id = ?", cartID).
		Preload("DualItemDiscount").
		Find(&appliedDualItemDiscount)
	appliedSingleItemCoupon := AppliedSingleItemCoupon{}
	db.Where("cart_id = ?", cartID).
		Preload("SingleItemCoupon").
		Find(&appliedSingleItemCoupon)

	for _, singleItemDiscount := range appliedSingleItemDiscount.SingleItemDiscount {
		if singleItemDiscount.Name != "" {
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
				Name:   singleItemDiscount.Name,
				Status: "APPLIED",
			})
		}
	}
	for _, dualItemDiscount := range appliedDualItemDiscount.DualItemDiscount {
		if dualItemDiscount.Name != "" {
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
				Name:   dualItemDiscount.Name,
				Status: "APPLIED",
			})
		}
	}
	for _, singeItemCoupon := range appliedSingleItemCoupon.SingleItemCoupon {
		if singeItemCoupon.Name != "" {
			fmt.Println("singeItemCoupon.Name", singeItemCoupon.Name)
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, Discount{
				Name:   singeItemCoupon.Name,
				Status: "APPLIED",
			})
		}
	}

	return &appliedDiscountsResponseList
}
