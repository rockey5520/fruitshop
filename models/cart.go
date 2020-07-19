package models

/*Cart is asssociated with Customer with Has-one relationship
swagger:model Cart
*/
type Cart struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the Cart table coming from the Customer table
	CustomerId uint `gorm:"not null;unique_index"`
	// Total amount valued for the cart
	Total float64 `json:"total"`
	// CartItem is having has-many relation with Cart
	CartItem []CartItem `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// Coupon is having has-many relation with Cart
	Coupon Coupon `gorm:"foreignkey:CartID;association_foreignkey:ID"`
	// Payment is having has-one relation with Cart
	Payment Payment `gorm:"foreignkey:CartID;association_foreignkey:ID"`
}
