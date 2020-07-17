package models

type Customer struct {
	ID        int    `json:"id" gorm:"primary_key"`
	LoginId   string `json:"loginid" gorm:"primary_key;unique_index"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
