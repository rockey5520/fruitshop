package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*
swagger:model Customer
*/
type Customer struct {
	// Login ID of the customer
	LoginID string `json:"loginid" gorm:"unique_index"`
	// First name of the customer
	FirstName string `json:"firstname"`
	// Last name of the customer
	LastName string `json:"lastname"`
	// Cart belongs to the customer and referenced with CustomerID as foreign key
	Cart Cart `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	gorm.Model
}

// SaveCustomer saves given customer record in Database
func (c *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {
	if err := db.Create(&c).Error; err != nil {
		return &Customer{}, err

	}
	return c, nil
}

//FindCustomerByLoginID provides information about a customer including its all realted tables
func (c *Customer) FindCustomerByLoginID(db *gorm.DB, loginID string) (*Customer, error) {
	err := db.Where("login_id = ?", loginID).First(&c).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer record Not Found")
	}
	// Preloading all the tables using cartID which joins all table, This reduces multiple lines of code and
	// represents simple and elegant for anyone reading code for first time
	cart := Cart{}
	db.Where("customer_id = ? AND status = ?", c.ID, "OPEN").
		Preload("CartItem").
		Preload("Payment").
		Preload("AppliedDualItemDiscount").
		Preload("AppliedSingleItemDiscount").
		Preload("AppliedSingleItemCoupon").
		Find(&cart)
	c.Cart = cart
	return c, err
}

// Validates given create customer payload
func (c *Customer) Validate(action string) error {

	if c.FirstName == "" {
		return errors.New("Required FirstName")
	}
	if c.LastName == "" {
		return errors.New("Required LastName")
	}
	return nil

}
