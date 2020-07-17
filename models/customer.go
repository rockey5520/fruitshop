package models

type Customer struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	LoginId   string `json:"loginid"`
}
