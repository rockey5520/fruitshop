package models

import "fruitshop/models"

type CartCustom struct {
	Cart     models.Cart
	CartItem []models.CartItem
	Fruit    []models.Fruit
}
