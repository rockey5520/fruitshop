package models

import "github.com/jinzhu/gorm"

/*
swagger:model AppliedSingleItemCoupon
*/
// AppliedSingleItemCoupon are active coupons applied per cart for coupon rules which are loaded to singleitemcoupon table
type AppliedSingleItemCoupon struct {
	// Primary key, created_at, deleted_at, updated_at for each applied single item coupon
	gorm.Model
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	//SingleItemCouponID is the primary key from the SingleItemCouponID table
	SingleItemCouponID uint
	//SingleItemCouponName  is the name of the coupon( Example : ORANGE30 )
	SingleItemCouponName string
	// Amount of the savings applied using this coupon
	Savings float64 `json:"savings"`
}
