package models

import "github.com/jinzhu/gorm"

type Customer struct {
	CustomerLoginId string     `json:"loginid" gorm:"primary_key;unique_index"`
	FirstName       string     `json:"firstname"`
	LastName        string     `json:"lastname"`
	Discounts       []Discount `gorm:"foreignkey:CustomerLoginId;association_foreignkey:CustomerLoginId"`
	gorm.Model
}
