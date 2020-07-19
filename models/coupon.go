package models

/*Coupon is asssociated with Cart with Has-many relationship
swagger:model Coupon
*/
type Coupon struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the Coupon table coming from the Cart table
	CartID uint `json:"cartid" gorm:"unique_index;not null"`
	// Name of the coupon
	Name string `json:"name" gorm:"not null;"`
	// Status of the coupon APPLIED and NOTAPPLIED are the two possible states
	Status string `json:"status"`
}
