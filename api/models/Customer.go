package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

/*Customer Customer table represents the customer of the fruit store
swagger:model Customer
*/
type Customer struct {
	// Login ID of the customer
	LoginID string `json:"loginid" gorm:"unique_index"`
	// First name of the customer
	FirstName string `json:"firstname"`
	// Last name of the customer
	LastName string `json:"lastname"`
	Cart     Cart   `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	gorm.Model
}

// HTTP status code 200 and Customer model in data
// swagger:response userResp
type swaggCustomerResp struct {
	// in:body
	Body struct {
		// HTTP status code 200
		Code int `json:"code"`
		// User model
		Data Customer `json:"data"`
	}
}

// Validate is a method
func (c *Customer) Validate(action string) error {

	if c.FirstName == "" {
		return errors.New("Required FirstName")
	}
	if c.LastName == "" {
		return errors.New("Required LastName")
	}
	return nil

}

// SaveCustomer is a
func (c *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	if err := db.Create(&c).Error; err != nil {
		return &Customer{}, err

	}
	return c, nil
}

//FindCustomerByID is a
func (c *Customer) FindCustomerByLoginID(db *gorm.DB, loginID string) (*Customer, error) {
	err := db.Where("login_id = ?", loginID).First(&c).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer record Not Found")
	}
	fmt.Println(c.FirstName)
	var cart Cart
	db.Where("customer_id = ? AND status = ?", c.ID, "OPEN").Find(&cart)
	c.Cart = cart
	// var cartItem []CartItem
	// db.Where("cart_id = ?", cart.ID).Find(&cartItem)
	// var payment Payment
	// db.Where("cart_id = ?", cart.ID).Find(&payment)
	// var appliedDualItemDiscount []AppliedDualItemDiscount
	// db.Where("cart_id = ?", cart.ID).Find(&appliedDualItemDiscount)
	// var appliedSingleItemDiscount []AppliedSingleItemDiscount
	// db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemDiscount)
	// var appliedSingleItemCoupon []AppliedSingleItemCoupon
	// db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemCoupon)
	// customer.Cart = cart
	// customer.Cart.CartItem = cartItem
	// customer.Cart.Payment = payment
	// customer.Cart.AppliedDualItemDiscount = appliedDualItemDiscount
	// customer.Cart.AppliedSingleItemCoupon = appliedSingleItemCoupon
	// customer.Cart.AppliedSingleItemDiscount = appliedSingleItemDiscount

	return c, err
}
