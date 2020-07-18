package models

type Cart struct {
	ID         uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CustomerId uint       `gorm:"not null;unique_index"`
	Total      float64    `json:"total"`
	CartItem   []CartItem `gorm:"foreignkey:CartID;association_foreignkey:ID"`
}
