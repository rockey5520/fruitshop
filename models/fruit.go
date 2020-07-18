package models

type Fruit struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name" gorm:"primary_key;unique_index"`
	Price float64 `json:"price"`
}
