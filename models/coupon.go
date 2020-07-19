package models

type Coupon struct {
	ID     uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CartID uint   `json:"cartid" gorm:"unique_index;not null"`
	Name   string `json:"name" gorm:"not null;"`
	Status string `json:"status"`
}
