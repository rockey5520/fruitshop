package models

import (
	"errors"

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

// H is a type swagger:response badReq
type H map[string]interface{}

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

	//var err error
	// update cart to cart array in the customer table
	newcart := Cart{
		Total:  0.0,
		Status: "OPEN",
	}

	customer := Customer{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		LoginID:   c.LoginID,
		Cart:      newcart,
	}

	if err := db.Create(&customer).Error; err != nil {
		return &Customer{}, err

	}
	return c, nil
}

//FindCustomerByID is a
func (c *Customer) FindCustomerByID(db *gorm.DB, loginID string) (*Customer, error) {
	var err error
	var customer Customer
	//db = c.MustGet("db").(*gorm.DB)
	err = db.Where("login_id = ?", loginID).First(&customer).Error

	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer record Not Found")
	}

	var cart Cart
	db.Where("customer_id = ? AND status = ?", customer.ID, "OPEN").Find(&cart)
	var cartItem []CartItem
	db.Where("cart_id = ?", cart.ID).Find(&cartItem)
	var payment Payment
	db.Where("cart_id = ?", cart.ID).Find(&payment)
	var appliedDualItemDiscount []AppliedDualItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedDualItemDiscount)
	var appliedSingleItemDiscount []AppliedSingleItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemDiscount)
	var appliedSingleItemCoupon []AppliedSingleItemCoupon
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemCoupon)
	customer.Cart = cart
	customer.Cart.CartItem = cartItem
	customer.Cart.Payment = payment
	customer.Cart.AppliedDualItemDiscount = appliedDualItemDiscount
	customer.Cart.AppliedSingleItemCoupon = appliedSingleItemCoupon
	customer.Cart.AppliedSingleItemDiscount = appliedSingleItemDiscount

	err = db.Debug().Model(Customer{}).Where("id = ?", loginID).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}

	return c, err
}
