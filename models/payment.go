package models

type Payment struct {
	ID     uint    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CartId uint    `gorm:"not null"`
	Amount float64 `json:"amount"`
	Status string  `json:"string"`
}
