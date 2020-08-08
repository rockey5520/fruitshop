package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*
swagger:model Payment
*/
// Payment is asssociated with Cart with Has-one relationship
type Payment struct {
	// Primary key for the Cart
	ID         uint `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CustomerID uint `json:"customerid"`
	// Foriegn key for the Payment table coming from the Cart table
	CartId uint `json:"cartid" gorm:"not null"`
	// Amount needs to paid by the customer for the purchase
	Amount float64 `json:"amount" binding:"required"`
	// Status of the coupon PAID and NOTPAID are the two possible states
	Status string `json:"string"`
}

// Pay method takes the payment and resets cart, cartitems, coupons, discounts
func (c *Payment) Pay(db *gorm.DB, payment Payment) (*Customer, error) {

	// Get cart
	cart := Cart{}
	db.Where("ID = ? AND status = ?", payment.CartId, "OPEN").Find(&cart)

	if cart.Total == payment.Amount && cart.Total != 0 && payment.Amount > 0 {
		// Empyt cart items table
		var cartItems []CartItem
		db.Find(&cartItems)

		// Set Cart amount to 0
		db.Model(&cart).Where("ID = ?", payment.CartId).Update("status", "CLOSED")

		pay := Payment{
			CustomerID: payment.CustomerID,
			CartId:     payment.CartId,
			Amount:     payment.Amount,
			Status:     "PAID",
		}
		// This creates creates the payment record in the table and closes the cart
		db.Create(&pay)

		newCart := Cart{
			CustomerId: payment.CustomerID,
			Total:      0.0,
			Status:     "OPEN",
		}
		// This creates creates the cart for the customer
		db.Create(&newCart)

	} else {
		return &Customer{}, errors.New("payment amount mismatched with the cart total")
	}

	var customer Customer
	err := db.Where("ID = ?", payment.CustomerID).Find(&customer).Error

	return &customer, err

}
