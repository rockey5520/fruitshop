package models

/*Discount is asssociated with Customer with Has-many relationship
swagger:model Discount
*/
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
