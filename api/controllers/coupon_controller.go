package controllers

import (
	"net/http"
	"strconv"
	"time"

	"fruitshop/api/models"
	"fruitshop/api/responses"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//ApplyTimeSensitiveCoupon applied coupon which is time sensitive based( in this use case its orange )
func (server *Server) ApplyTimeSensitiveCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cart_id := vars["cart_id"]
	fruit_id := vars["fruit_id"]

	//ApplyTimeSensitiveCoupon(server.DB, cart_id, fruit_id)
	go ApplySingleItemTimSensitiveCoupon(server.DB, cart_id, fruit_id)

	responses.JSON(w, http.StatusOK, "Applied discount")
}

//ApplyOrangeCoupon applies discount for oranges
func ApplyTimeSensitiveCoupon(db *gorm.DB, cart_id string, fruit_id string) {
	// This coupon will be run as go routine and sleeps for 10 seconds

}

// ApplySingleItemTimSensitiveCoupon is
func ApplySingleItemTimSensitiveCoupon(db *gorm.DB, cart_id string, fruit_id string) {

	fruit := models.Fruit{}
	db.Where("ID = ?", fruit_id).
		Preload("SingleItemCoupon").
		Find(&fruit)

	var singleItemCouponList []models.SingleItemCoupon
	db.Where("fruit_id = ?", fruit_id).Find(&singleItemCouponList)
	var interestCoupon models.SingleItemCoupon
	for _, coupon := range singleItemCouponList {
		if coupon.FruitID == fruit.ID {
			interestCoupon = coupon
		}
	}
	cartID, _ := strconv.Atoi(cart_id)
	appliedSingleItemCoupon := models.AppliedSingleItemCoupon{
		CartID:           uint(cartID),
		SingleItemCoupon: singleItemCouponList,
	}

	var cartItem models.CartItem
	db.Where("cart_id = ? AND fruit_id = ?", cart_id, fruit_id).Find(&cartItem)
	if cartItem.Quantity > 0 {
		discountCalculated := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(interestCoupon.Discount)
		updatedTotalCost := cartItem.ItemTotal - discountCalculated
		db.Model(&cartItem).
			Where("cart_id = ?", cartItem.CartID).
			Update("ItemTotal", updatedTotalCost).
			Update("item_discounted_total", discountCalculated)
		RecalcualtePayments(db, uint(cartID))
		appliedSingleItemCoupon.Savings = discountCalculated
		if err := db.Model(&appliedSingleItemCoupon).
			Where("cart_id = ?", cartItem.CartID).
			First(&appliedSingleItemCoupon).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&appliedSingleItemCoupon)
			}
		} else {
			db.Model(&appliedSingleItemCoupon).
				Where("cart_id = ? ", cart_id).
				Update("savings", discountCalculated)
		}
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	var cart models.Cart
	db.Where("ID = ?", cart_id).Find(&cart)
	if cart.Status != "CLOSED" {
		var cartItem models.CartItem
		db.Where("ID = ?", cart_id).First((&cartItem))
		var fruit models.Fruit
		db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		db.Model(&cartItem).
			Where("cart_id = ?", cartItem.ID).
			Update("ItemTotal", float64(cartItem.Quantity)*fruit.Price).
			Update("item_discounted_total", 0.0)
		RecalcualtePayments(db, cartItem.CartID)
		db.Unscoped().Where("cart_id = ?", cart_id).Delete(&appliedSingleItemCoupon)
	}

}
