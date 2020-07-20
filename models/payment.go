package models

/*Payment is asssociated with Caet with Has-one relationship
swagger:model Payment
*/
type Payment struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the Payment table coming from the Cart table
	uint `gorm:"not null"`
	// Amount needs to paid by the customer for the purchase
	Amount float64 `json:"amount"`
	// Status of the coupon PAID and NOTPAID are the two possible states
	Status string `json:"string"`
}
