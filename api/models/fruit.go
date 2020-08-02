package models

import "github.com/jinzhu/gorm"

/*Fruit is a meta table of fruits in the inventory
swagger:model Fruit
*/
// Fruit is
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

// FindAllFruits is
func (f *Fruit) FindAllFruits(db *gorm.DB) (*[]Fruit, error) {
	var err error
	fruits := []Fruit{}
	err = db.Debug().Model(&Fruit{}).Limit(100).Find(&fruits).Error
	if err != nil {
		return &[]Fruit{}, err
	}
	return &fruits, err
}
