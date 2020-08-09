package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

/*
swagger:model CartItem
*/
// CartItem represents struct of the cartitem added to a cart
type CartItem struct {
	// Primary key for the Cart
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	// Foriegn key for the CartItem table coming from the Cart table
	CartID uint `gorm:"not null"`
	// FruitID is represents the ID of the fruit added to cart
	FruitID uint `gorm:"not null"`
	// Name of the fruit
	Name string `json:"name" gorm:"not null"`
	// Number of fruits ordered
	Quantity int `json:"quantity"`
	// Total cost for this fruits based on number of items
	ItemTotal float64 `json:"itemtotal"`
	// Total discounted cost for this fruits based on number of items
	ItemDiscountedTotal float64 `json:"ItemDiscountedTotal"`
}

// SaveOrUpdateCartItem saves the cart item entry to the DB if given quantity is > 0 or removes if quantity is 0
func (input *CartItem) SaveCartItem(db *gorm.DB) (*CartItem, error) {
	var err error
	var fruit Fruit
	db.Where("name = ?", input.Name).First(&fruit)

	input.ItemTotal = fruit.Price * float64(input.Quantity)
	input.ItemDiscountedTotal = 0.0
	input.FruitID = fruit.ID

	cartItem := CartItem{
		CartID:              input.CartID,
		FruitID:             fruit.ID,
		Name:                input.Name,
		ItemTotal:           fruit.Price * float64(input.Quantity),
		ItemDiscountedTotal: 0.0,
		Quantity:            input.Quantity,
	}

	if input.Quantity > 0 {
		// Create/update fruit to the cart
		if err = db.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", input.CartID, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&cartItem) // create new record from newUser
			}
		}
	}

	return &cartItem, nil
}

// UpdateCartItem saves the cart item entry to the DB if given quantity is > 0 or removes if quantity is 0
func (input *CartItem) UpdateCartItem(db *gorm.DB) (*CartItem, error) {
	var err error
	var fruit Fruit
	db.Where("name = ?", input.Name).First(&fruit)

	input.ItemTotal = fruit.Price * float64(input.Quantity)
	input.ItemDiscountedTotal = 0.0
	input.FruitID = fruit.ID

	cartItem := CartItem{
		CartID:              input.CartID,
		FruitID:             fruit.ID,
		Name:                input.Name,
		ItemTotal:           fruit.Price * float64(input.Quantity),
		ItemDiscountedTotal: 0.0,
		Quantity:            input.Quantity,
	}

	if input.Quantity > 0 {
		// update fruit to the cart
		if err = db.Model(&cartItem).Where("cart_id = ? AND fruit_id = ? ", input.CartID, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&cartItem) // create new record from newUser
			}
		} else {
			err = db.Model(&cartItem).Where("cart_id = ?  AND fruit_id = ? ", input.CartID, fruit.ID).
				Update("quantity", input.Quantity).
				Update("fruit_id", fruit.ID).
				Update("item_total", float64(input.Quantity)*fruit.Price).
				Update("item_discounted_total", 0.0).Error
		}

	}
	return &cartItem, err
}

// DeleteCartItem saves the cart item entry to the DB if given quantity is > 0 or removes if quantity is 0
func (input *CartItem) DeleteCartItem(db *gorm.DB, cartid string, fruitName string) (*CartItem, error) {
	var err error
	cartItem := CartItem{}

	if input.Quantity == 0 {
		err = db.Where("cart_id = ? AND name = ?", cartid, fruitName).Delete(&cartItem).Error

	}
	return &cartItem, err
}

// FindAllCartItems returns all items present in a particular cart using cartID
func (c *CartItem) FindAllCartItems(db *gorm.DB, cartID string) *[]CartItemResponse {
	cartItemsArray := make([]CartItemResponse, 0)

	var cartItems []CartItem
	db.Where("cart_id = ?", cartID).Find(&cartItems)

	for _, cartItem := range cartItems {
		var fruit Fruit
		db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		cartItemsArray = append(cartItemsArray, CartItemResponse{
			Name:        fruit.Name,
			CostPerItem: fruit.Price,
			Count:       cartItem.Quantity,
			ItemTotal:   cartItem.ItemTotal,
		})

	}
	return &cartItemsArray
}

// Validates given cartitem payload
func (c *CartItem) Validate(action string) error {
	if c.Name == "" {
		return errors.New("Required name")
	}
	if c.Quantity < 0 {
		return errors.New("Required valid Quantity")
	}
	return nil
}

// RecalcualtePayments recalcuates the cart value and its saving based on the cart items
func RecalcualtePayments(db *gorm.DB, cartID uint) {
	// Fetch all items in a given cart
	var cartItems []CartItem
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}

	// calcuate the total cost of the cartitems and total discounts applied
	var totalCost float64
	var totalDiscountedCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
		totalDiscountedCost += item.ItemDiscountedTotal
	}
	var cart Cart
	if err := db.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}

	// Update Cart table with total cost and total savings
	db.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
}

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
