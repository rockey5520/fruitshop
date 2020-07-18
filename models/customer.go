package models

import "github.com/jinzhu/gorm"

type Customer struct {
	LoginId   string `json:"loginid" gorm:"unique_index"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	//Discounts       []Discount `gorm:"foreignkey:CustomerLoginId;association_foreignkey:CustomerLoginId"`
	Discounts []Discount `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	Cart      Cart       `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`
	gorm.Model
}
