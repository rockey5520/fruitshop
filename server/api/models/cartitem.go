package models

import (
	"github.com/jinzhu/gorm"
)

/*CartItem is asssociated with Cart with Has-many relationship
swagger:model CartItem
*/
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
	// Total disQuantityed cost for this fruits based on number of items
	ItemDisQuantityedTotal float64 `json:"ItemDisQuantityedTotal"`
}

//SaveUpdateCartItem is
func (c *CartItem) SaveUpdateCartItem(db *gorm.DB) (*CartItem, error) {
	//var err error
	//db = cartItem.MustGet("db").(*gorm.DB)
	// Bind the input payload to schema for validations
	// var input CartItemInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	var fruit Fruit
	db.Where("name = ?", c.FruitID).First(&fruit)

	//Create/Update/Delete Cart entry based on the Quantity
	cartItem := CartItem{CartID: c.CartID, FruitID: fruit.ID, ItemTotal: fruit.Price * float64(c.Quantity), ItemDisQuantityedTotal: 0.0}
	if c.Quantity > 0 {
		// Create/update fruit to the cart
		cartItem.Quantity = c.Quantity
		if err := db.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", c.CartID, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&cartItem) // create new record from newUser
			}
		} else {
			db.Model(&cartItem).Where("cart_id = ?  AND fruit_id = ? ", c.CartID, fruit.ID).
				Update("quantity", c.Quantity).
				Update("fruit_id", fruit.ID).
				Update("item_total", float64(c.Quantity)*fruit.Price).
				Update("item_disQuantityed_total", 0.0)
		}
	} else if c.Quantity == 0 {
		db.Where("cart_id = ? AND fruit_id = ?", c.CartID, fruit.ID).Delete(&cartItem)

	}
	return c, nil
}