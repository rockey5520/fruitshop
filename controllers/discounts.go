package controllers

import (
	"fmt"
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
	appliedDiscountsResponseList := make([]models.AppliedDiscountsResponse, 0)

	var appliedSingleItemDiscount models.AppliedSingleItemDiscount
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&appliedSingleItemDiscount)
	var singleItemDiscount models.SingleItemDiscount
	models.DB.Where("ID = ?", appliedSingleItemDiscount.SingleItemDiscountID).Find(&singleItemDiscount)

	var appliedDualItemDiscount models.AppliedDualItemDiscount
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&appliedDualItemDiscount)
	var dualItemDiscount models.DualItemDiscount
	models.DB.Where("ID = ?", appliedDualItemDiscount.DualItemDiscountID).Find(&dualItemDiscount)

	var appliedSingleItemCoupon models.AppliedSingleItemCoupon
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&appliedSingleItemCoupon)
	var singeItemCoupon models.SingleItemCoupon
	models.DB.Where("ID = ?", appliedSingleItemCoupon.SingleItemCouponID).Find(&singeItemCoupon)

	if singleItemDiscount.Name != "" {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
			Name:   singleItemDiscount.Name,
			Status: "APPLIED",
		})
	}
	if dualItemDiscount.Name != "" {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
			Name:   dualItemDiscount.Name,
			Status: "APPLIED",
		})
	}
	if singeItemCoupon.Name != "" {
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
			Name:   singeItemCoupon.Name,
			Status: "APPLIED",
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": appliedDiscountsResponseList})
}

// @Summary Applied orange 30 coupon code
// @Description This endpoint applied orange 30 percent discount coupon code
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"

// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/orangecoupon/{cart_id} [get]
//ApplyOrangeCoupon applies discount for oranges
func ApplyOrangeCoupon(c *gin.Context) {
	// This coupon will be run as go routine and sleeps for 30 seconds
	go ApplySingleItemCoupon(c, "ORANGE30")
}

func ApplySingleItemCoupon(c *gin.Context, discountCouponCode string) {

	var cartItems []models.CartItem
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&cartItems)
	var singeItemCoupon models.SingleItemCoupon
	models.DB.Where("name = ?", discountCouponCode).Find(&singeItemCoupon)
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
			RecalcualtePayments(cartItem.CartID)
			appliedItemCoupon.Savings = discountCalculated
			if err := models.DB.Model(&appliedItemCoupon).
				Where("cart_id = ?", cartItem.CartID).
				First(&appliedItemCoupon).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					models.DB.Create(&appliedItemCoupon)
				}
			} else {
				models.DB.Model(&appliedItemCoupon).Where("cart_id = ? ", c.Param("cart_id")).Update("savings", discountCalculated)
			}
		}
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	var cart models.Cart
	models.DB.Where("ID = ?", c.Param("cart_id")).Find(&cart)
	if cart.Status != "CLOSED" {
		var cartItem models.CartItem
		models.DB.Where("ID = ?", c.Param("cart_id")).First((&cartItem))
		var fruit models.Fruit
		models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.ID).Update("ItemTotal", float64(cartItem.Quantity)*fruit.Price)
		RecalcualtePayments(cartItem.CartID)
		models.DB.Unscoped().Where("cart_id = ?", c.Param("cart_id")).Delete(&appliedItemCoupon)
	}

}

// ApplySingleItemDiscounts applies single fruit discounts based on the single item discounts
func ApplySingleItemDiscounts(cartItem models.CartItem) {

	// This applies 10 percent discount for APPLE and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplySingleItemDiscount(cartItem, "APPLE10")

}

// ApplyDualItemDiscounts applies discount for dual items discounts based on the dual item combo  discounts
func ApplyDualItemDiscounts(cartItem models.CartItem) {

	// This applies 30% discount for Banana pear set and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplyDualItemDiscount(cartItem, "PEARBANANA30")

}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyDualItemDiscount(cartItem models.CartItem, discountCouponCode string) {
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount models.DualItemDiscount
	models.DB.Where("name = ?", discountCouponCode).Find(&dualItemDiscount)

	appliedDualItemDiscount := models.AppliedDualItemDiscount{
		CartID:             cartItem.CartID,
		DualItemDiscountID: dualItemDiscount.ID,
	}

	var fruitId1Count int
	var fruitId2Count int
	var discountUpdate bool

	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cartItem.CartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	for _, x := range cartItems {
		cartItem := x
		if cartItem.FruitID == dualItemDiscount.FruitID_1 {
			fruitId1Count = cartItem.Quantity

		} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
			fruitId2Count = cartItem.Quantity
		}
	}

	sets := getSets(fruitId1Count, fruitId2Count, dualItemDiscount.Count_1, dualItemDiscount.Count_2)

	if sets != 0 {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.FruitID == dualItemDiscount.FruitID_1 {
				var fruit models.Fruit
				models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_1) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings += discount
				models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount)
			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				var fruit models.Fruit
				models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings += discount
				models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount)
			}
		}
		discountUpdate = true
	} else {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.FruitID == dualItemDiscount.FruitID_1 {
				models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity) * fruit.Price))
			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity) * fruit.Price))
			}
		}

	}

	if discountUpdate {
		if err := models.DB.Model(&appliedDualItemDiscount).
			Where("cart_id = ?", cartItem.CartID).
			First(&appliedDualItemDiscount).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&appliedDualItemDiscount) // create new record from newUser
			}
		}
	} else {
		models.DB.Unscoped().Where("cart_id = ?", cartItem.CartID).Delete(&appliedDualItemDiscount)
	}

}

//ApplyApple10Discount applies apple 10 percent discount
func ApplySingleItemDiscount(cartItem models.CartItem, discountCouponCode string) {
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var singeItemDiscount models.SingleItemDiscount
	models.DB.Where("fruit_id = ?", fruit.ID).Find(&singeItemDiscount)
	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{
		CartID:               cartItem.CartID,
		SingleItemDiscountID: singeItemDiscount.ID,
	}
	if cartItem.FruitID == fruit.ID && singeItemDiscount.Name == discountCouponCode {
		if cartItem.Quantity >= 7 {
			var itemCost float64
			var appleCost float64

			appleCost += float64(cartItem.Quantity) * fruit.Price

			discount := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(singeItemDiscount.Discount)
			itemCost = (appleCost - discount)
			appliedSingleItemDiscount.Savings = discount
			models.DB.Model(&cartItem).Update("item_total", itemCost)

			if err := models.DB.Model(&appliedSingleItemDiscount).
				Where("cart_id = ?", cartItem.CartID).
				First(&appliedSingleItemDiscount).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					models.DB.Create(&appliedSingleItemDiscount) // create new record from newUser
				}
			}
		} else {
			var itemCost float64
			itemCost += float64(cartItem.Quantity) * fruit.Price
			models.DB.Model(&cartItem).Update("item_total", itemCost)
			models.DB.Unscoped().Where("cart_id = ?", cartItem.CartID).Delete(&appliedSingleItemDiscount)
		}
	}
}

// getSets gets the number of sets of banana and pears
func getSets(fruit1 int, fruit2 int, count1 int, count2 int) int {

	fruit1Count := fruit1
	fruit2Count := fruit2
	var set int

	for i := 0; i < fruit1Count; i++ {
		if fruit1Count >= count1 && fruit2Count >= count2 {
			set += 1
			fruit1Count = fruit1Count - count1
			fruit2Count = fruit2Count - count2
		} else {
			break
		}
	}

	fmt.Println("set ", set)
	return set
}
