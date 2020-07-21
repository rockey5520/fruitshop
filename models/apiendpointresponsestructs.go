package models

// Creates a customer details row in the database
//swagger:parameters CreateCustomerInput
type CreateCustomerInput struct {
	// First name of the customer
	FirstName string `json:"firstname" binding:"required"`
	// Last name of the customer
	LastName string `json:"lastname" binding:"required"`
	// Login ID of the customer
	LoginId string `json:"loginid" binding:"required"`
}

type DiscountsApplied struct {
	DiscountNames []string `json:"discounts" binding:"required"`
}
