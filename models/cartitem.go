package models

/*CartItem is asssociated with Cart with Has-many relationship
swagger:model CartItem
*/
type CartItem struct {
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
