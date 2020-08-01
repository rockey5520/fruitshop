package models

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
	// Total discounted cost for this fruits based on number of items
	ItemDiscountedTotal float64 `json:"ItemDiscountedTotal"`
}

func (cartItem *CartItem) SaveUpdateCartItem(db *gorm.DB) (*CartItem, error) {
	var err error
	db := cartItem.MustGet("db").(*gorm.DB)
	// Bind the input payload to schema for validations
	var input CartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fruit models.Fruit
	if err := db.Where("name = ?", input.Name).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: input.CartId, FruitID: fruit.ID, ItemTotal: fruit.Price * float64(input.Count), ItemDiscountedTotal: 0.0}
	if input.Count > 0 {
		// Create/update fruit to the cart
		cartItem.Quantity = input.Count
		if err := db.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", input.CartId, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&cartItem) // create new record from newUser
			}
		} else {
			db.Model(&cartItem).Where("cart_id = ?  AND fruit_id = ? ", input.CartId, fruit.ID).
				Update("quantity", input.Count).
				Update("fruit_id", fruit.ID).
				Update("item_total", float64(input.Count)*fruit.Price).
				Update("item_discounted_total", 0.0)
		}
	} else if input.Count == 0 {
		db.Where("cart_id = ? AND fruit_id = ?", input.CartId, fruit.ID).Delete(&cartItem)

	}
	return cartItem, nil
}
