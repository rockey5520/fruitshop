package models

import "github.com/jinzhu/gorm"

/*Customer Customer table represents the customer of the fruit store
swagger:model Customer
*/
type Customer struct {
	// Login ID of the customer
	LoginId string `json:"loginid" gorm:"unique_index"`
	// First name of the customer
	FirstName string `json:"firstname"`
	// Last name of the customer
	LastName string `json:"lastname"`
	// Discounts assosiated to the customer
	Discounts []Discount `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	// Cart assosiated to the customer
	Cart Cart `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	
	gorm.Model
}

// Create a new product
//swagger:parameters CreateCustomerInput
type CreateCustomerInput struct {
	// First name of the customer
	FirstName string `json:"firstname" binding:"required"`
	// Last name of the customer
	LastName string `json:"lastname" binding:"required"`
	// Login ID of the customer
	LoginId string `json:"loginid" binding:"required"`
}

// Error Bad Request
// swagger:response badReq
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
