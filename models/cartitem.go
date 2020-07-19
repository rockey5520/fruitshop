package models

type CartItem struct {
	ID          uint    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CartID      uint    `gorm:"not null"`
	Name        string  `json:"name" gorm:"not null;"`
	CostPerItem float64 `json:"costperitem" gorm:"not null;"`
	Count       int     `json:"count"`
	ItemTotal   float64 `json:"itemtotal"`
}
