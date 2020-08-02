package models

/*
CartItem is asssociated with Cart with Has-many relationship
swagger:model CartItem
*/
// CartItem is
type CartItem struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// Fruit identifier
	//Fruit Fruit `gorm:"foreignkey:ID;association_foreignkey:ID"`
	FruitID uint `gorm:"not null"`
	// Number of fruits ordered
	Quantity int `json:"quantity"`
	// Total cost for this fruits based on number of items
	ItemTotal float64 `json:"itemtotal"`
	// Total discounted cost for this fruits based on number of items
	ItemDiscountedTotal float64 `json:"ItemDiscountedTotal"`
}
