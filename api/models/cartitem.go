package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

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
	//Fruit Fruit `gorm:"foreignkey:ID;association_foreignkey:ID"`
	Name string `json:"name" gorm:"not null"`
	// Number of fruits ordered
	Quantity int `json:"quantity"`
	// Total cost for this fruits based on number of items
	ItemTotal float64 `json:"itemtotal"`
	// Total discounted cost for this fruits based on number of items
	ItemDiscountedTotal float64 `json:"ItemDiscountedTotal"`
}

// Validate is a method
func (c *CartItem) Validate(action string) error {

	if c.Name == "" {
		return errors.New("Required name")
	}
	if c.Quantity < 0 {
		return errors.New("Required valid Quantity")
	}
	return nil

}

// SaveOrUpdateCartItem is a
func (input *CartItem) SaveOrUpdateCartItem(db *gorm.DB) (*CartItem, error) {

	var fruit Fruit
	db.Where("name = ?", input.Name).First(&fruit)

	input.ItemTotal = fruit.Price * float64(input.Quantity)
	input.ItemDiscountedTotal = 0.0
	input.FruitID = fruit.ID

	cartItem := CartItem{
		CartID:              input.CartID,
		FruitID:             fruit.ID,
		ItemTotal:           fruit.Price * float64(input.Quantity),
		ItemDiscountedTotal: 0.0,
		Quantity:            input.Quantity,
	}

	if input.Quantity > 0 {
		// Create/update fruit to the cart
		if err := db.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", input.CartID, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&cartItem) // create new record from newUser
			}
		} else {
			fmt.Println(input.Quantity)
			db.Model(&cartItem).Where("cart_id = ?  AND fruit_id = ? ", input.CartID, fruit.ID).
				Update("quantity", input.Quantity).
				Update("fruit_id", fruit.ID).
				Update("item_total", float64(input.Quantity)*fruit.Price).
				Update("item_discounted_total", 0.0)
		}
	} else if input.Quantity == 0 {
		db.Where("cart_id = ? AND fruit_id = ?", input.CartID, fruit.ID).Delete(&cartItem)

	}
	return &cartItem, nil
}
