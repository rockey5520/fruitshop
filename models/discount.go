package models

type Discount struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CustomerId uint   `gorm:"not null"`
	Name       string `json:"name"`
	Status     string `json:"string"`
}
