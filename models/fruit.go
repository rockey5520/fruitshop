package models

/*Fruit is a meta table of fruits in the inventory
swagger:model Fruit
*/
type Fruit struct {
	// Primary key for the Cart
	ID int `json:"id" gorm:"primary_key"`
	// Name of the fruit
	Name string `json:"name" gorm:"primary_key;unique_index"`
	// Price of each fruit
	Price float64 `json:"price"`
}
