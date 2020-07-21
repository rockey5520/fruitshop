package controllers

import (
	"fruitshop/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary Show details of all discounts for a given cart
// @Description Get details of all discounts for each cart
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"
// @Success 200 {array} models.Discount
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/discounts/{cart_id} [get]

// FindDiscounts will return all discounts with status as APPLIED available within the fruitshop
func FindDiscounts(c *gin.Context) {
	s := make([]string, 0)

	var appliedSingleItemDiscount models.AppliedSingleItemDiscount
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&appliedSingleItemDiscount)
	var singleItemDiscount models.SingleItemDiscount
	models.DB.Where("ID = ?", appliedSingleItemDiscount.SingleItemDiscountID).Find(&singleItemDiscount)

	var appliedDualItemDiscount models.AppliedDualItemDiscount
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&appliedDualItemDiscount)
	var dualItemDiscount models.DualItemDiscount
	models.DB.Where("ID = ?", appliedDualItemDiscount.DualItemDiscountID).Find(&dualItemDiscount)

	s = append(s, singleItemDiscount.Name, dualItemDiscount.Name)

	discountsApplied := models.DiscountsApplied{
		DiscountNames: s,
	}

	c.JSON(http.StatusOK, gin.H{"data": discountsApplied})
}

// @Summary Applied orange 30 coupon code
// @Description This endpoint applied orange 30 percent discount coupon code
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"
// @Success 200 {object} models.Coupon
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/orangecoupon/{cart_id} [get]
//ApplyOrangeCoupon applies discount for oranges
func ApplyOrangeCoupon(c *gin.Context) {
	go ApplyOrange30Coupon(c)
}

func ApplyOrange30Coupon(c *gin.Context) {

	var cartItems []models.CartItem
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&cartItems)
	var singeItemCoupon models.SingleItemCoupon
	models.DB.Where("name = ?", "ORANGE30").Find(&singeItemCoupon)
	appliedItemCoupon := models.AppliedSingleItemCoupon{
		CartID:             cartItems[0].CartID,
		SingleItemCouponID: singeItemCoupon.ID,
	}

	for _, cartItem := range cartItems {
		var fruit models.Fruit
		models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		if fruit.Name == "Orange" && cartItem.Quantity > 0 {
			var singleItemCoupon models.SingleItemCoupon
			models.DB.Where("fruit_id = ?", fruit.ID).Find(&singleItemCoupon)
			discountCalculated := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(singleItemCoupon.Discount)
			updatedTotalCost := cartItem.ItemTotal - discountCalculated
			models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("ItemTotal", updatedTotalCost)
			appliedItemCoupon.Savings = discountCalculated
			if err := models.DB.Model(&appliedItemCoupon).
				Where("cart_id = ?", cartItem.CartID).
				First(&appliedItemCoupon).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					models.DB.Create(&appliedItemCoupon) // create new record from newUser
				}
			} else {
				models.DB.Model(&appliedItemCoupon).Where("cart_id = ? ", c.Param("cart_id")).Update("savings", discountCalculated)
			}
		}
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	var cartItem models.CartItem
	models.DB.Where("ID = ?", c.Param("cart_id")).FirstOrInit((&cartItem))
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.ID).Update("ItemTotal", float64(cartItem.Quantity)*fruit.Price)
	models.DB.Unscoped().Where("cart_id = ?", c.Param("cart_id")).Delete(&appliedItemCoupon)

	RecalcualtePayments(cartItem.CartID)

}
