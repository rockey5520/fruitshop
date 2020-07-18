package models

type Discount struct {
	ID              uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	CustomerLoginId string `gorm:"not null"`
	Name            string `json:"name" gorm:"unique_index"`
	Status          string `json:"string"`
}
