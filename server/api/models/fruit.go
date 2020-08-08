package models

import "github.com/jinzhu/gorm"

/*
swagger:model Fruit
*/
// Fruit is a meta table of fruits in the inventory
type Fruit struct {
	// Primary key, created_at, deleted_at, updated_at for each fruit
	gorm.Model
	// Name of the fruit
	Name string `json:"name" gorm:"primary_key;unique_index"`
	// Price of each fruit
	Price float64 `json:"price"`
	// Single Item Discount
	SingleItemDiscount []SingleItemDiscount `gorm:"foreignkey:FruitID"`
	// Single Item Coupon
	SingleItemCoupon []SingleItemCoupon

	DualItemDiscount []DualItemDiscount `gorm:"foreignkey:FruitID"`
}

// FindAllFruits fetched all fruits available in the database
func (f *Fruit) FindAllFruits(db *gorm.DB) (*[]Fruit, error) {
	var err error
	fruits := []Fruit{}
	err = db.Debug().Model(&Fruit{}).Limit(4).Find(&fruits).Error
	if err != nil {
		return &[]Fruit{}, err
	}
	return &fruits, err
}
