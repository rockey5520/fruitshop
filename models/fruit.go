package models

type Fruit struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name" gorm:"primary_key;unique_index"`
	Price float64 `json:"price"`
}

type Customer1 struct {
	ID        int    `json:"id" gorm:"primary_key"`
	LoginId   string `json:"loginid" gorm:"primary_key;unique_index"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
