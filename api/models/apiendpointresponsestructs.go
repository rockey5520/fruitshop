package models

// Creates a customer details row in the database
//swagger:parameters CreateCustomerInput
type CreateCustomerInput struct {
	// First name of the customer
	FirstName string `json:"firstname" binding:"required"`
	// Last name of the customer
	LastName string `json:"lastname" binding:"required"`
	// Login ID of the customer
	LoginId string `json:"loginid" binding:"required"`
}

/*CartItemResponse is asssociated with Cart with Has-many relationship
swagger:model CartItemResponse
*/
type CartItemResponse struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// Name of the Fruit
	Name string `json:"name" gorm:"not null;"`
	// Cost per fruit
	CostPerItem float64 `json:"costperitem" gorm:"not null;"`
	// Number of fruits ordered
	Count int `json:"count"`
	// Total cost for this fruits based on number of items
	ItemTotal float64 `json:"itemtotal"`
}
