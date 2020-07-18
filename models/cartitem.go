package models

import "github.com/jinzhu/gorm"

type CartItem struct {
	//ID              uint    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CartID uint   `gorm:"not null"`
	Name   string `json:"name" gorm:"not null;"`
	Count  int    `json:"count"`
	gorm.Model
}
