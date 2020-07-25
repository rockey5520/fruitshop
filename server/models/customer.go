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
	Cart     Cart   `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`

	gorm.Model
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
