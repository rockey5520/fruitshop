package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"
	"strconv"
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

	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{}
	models.DB.Where("cart_id = ?", c.Param("cart_id")).
		Preload("SingleItemDiscount").
		Find(&appliedSingleItemDiscount)
	appliedDualItemDiscount := models.AppliedDualItemDiscount{}
	models.DB.Where("cart_id = ?", c.Param("cart_id")).
		Preload("DualItemDiscount").
		Find(&appliedDualItemDiscount)
	appliedSingleItemCoupon := models.AppliedSingleItemCoupon{}
	models.DB.Where("cart_id = ?", c.Param("cart_id")).
		Preload("SingleItemCoupon").
		Find(&appliedSingleItemCoupon)

	for _, singleItemDiscount := range appliedSingleItemDiscount.SingleItemDiscount {
		if singleItemDiscount.Name != "" {
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
				Name:   singleItemDiscount.Name,
				Status: "APPLIED",
			})
		}
	}
	for _, dualItemDiscount := range appliedDualItemDiscount.DualItemDiscount {
		if dualItemDiscount.Name != "" {
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
				Name:   dualItemDiscount.Name,
				Status: "APPLIED",
			})
		}
	}
	for _, singeItemCoupon := range appliedSingleItemCoupon.SingleItemCoupon {
		if singeItemCoupon.Name != "" {
			fmt.Println("singeItemCoupon.Name", singeItemCoupon.Name)
			appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
				Name:   singeItemCoupon.Name,
				Status: "APPLIED",
			})
		}
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
func ApplyTimeSensitiveCoupon(c *gin.Context) {
	// This coupon will be run as go routine and sleeps for 10 seconds
	go ApplySingleItemTimSensitiveCoupon(c)
}

func ApplySingleItemTimSensitiveCoupon(c *gin.Context) {
	cart_id, _ := strconv.Atoi(c.Param("cart_id"))
	fruit_id, _ := strconv.Atoi(c.Param("fruit_id"))

	fruit := models.Fruit{}
	models.DB.Where("ID = ?", fruit_id).
		Preload("SingleItemCoupon").
		Find(&fruit)

	//singleItemCouponList := fruit.SingleItemCoupon
	var singleItemCouponList []models.SingleItemCoupon
	models.DB.Where("fruit_id = ?", fruit_id).Find(&singleItemCouponList)
	var interestCoupon models.SingleItemCoupon
	for _, coupon := range singleItemCouponList {
		if coupon.FruitID == fruit.ID {
			interestCoupon = coupon
		}
	}

	appliedSingleItemCoupon := models.AppliedSingleItemCoupon{
		CartID:           uint(cart_id),
		SingleItemCoupon: singleItemCouponList,
	}

	var cartItem models.CartItem
	models.DB.Where("cart_id = ? AND fruit_id = ?", cart_id, fruit_id).Find(&cartItem)

	/* var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit) */
	if cartItem.Quantity > 0 {
		discountCalculated := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(interestCoupon.Discount)
		updatedTotalCost := cartItem.ItemTotal - discountCalculated
		models.DB.Model(&cartItem).
			Where("cart_id = ?", cartItem.CartID).
			Update("ItemTotal", updatedTotalCost).
			Update("item_discounted_total", discountCalculated)
		RecalcualtePayments(cartItem.CartID)
		appliedSingleItemCoupon.Savings = discountCalculated
		if err := models.DB.Model(&appliedSingleItemCoupon).
			Where("cart_id = ?", cartItem.CartID).
			First(&appliedSingleItemCoupon).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&appliedSingleItemCoupon)
			}
		} else {
			models.DB.Model(&appliedSingleItemCoupon).
				Where("cart_id = ? ", c.Param("cart_id")).
				Update("savings", discountCalculated)
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
		models.DB.Model(&cartItem).
			Where("cart_id = ?", cartItem.ID).
			Update("ItemTotal", float64(cartItem.Quantity)*fruit.Price).
			Update("item_discounted_total", 0.0)
		RecalcualtePayments(cartItem.CartID)
		models.DB.Unscoped().Where("cart_id = ?", c.Param("cart_id")).Delete(&appliedSingleItemCoupon)
	}

}

// ApplySingleItemDiscounts applies single fruit discounts based on the single item discounts
func ApplySingleItemDiscounts(cartItem models.CartItem) {

	// This applies 10 percent discount for APPLE and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplySingleItemDiscount(cartItem)

}

// ApplyDualItemDiscounts applies discount for dual items discounts based on the dual item combo  discounts
func ApplyDualItemDiscounts(cartItem models.CartItem) {

	// This applies 30% discount for Banana pear set and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplyDualItemDiscount(cartItem)

}

//ApplyApple10Discount applies apple 10 percent discount
func ApplySingleItemDiscount(cartItem models.CartItem) {
	fruit := models.Fruit{}
	models.DB.Where("ID = ?", cartItem.FruitID).
		Preload("SingleItemDiscount").
		Find(&fruit)

	singeItemDiscount := fruit.SingleItemDiscount
	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{
		CartID:             cartItem.CartID,
		SingleItemDiscount: singeItemDiscount,
	}

	for _, singleItemDiscount := range singeItemDiscount {
		if cartItem.FruitID == singleItemDiscount.FruitID {
			fmt.Println("singleItemDiscount.Count", singleItemDiscount.Count)
			if cartItem.Quantity >= singleItemDiscount.Count {
				var calculatedCost float64
				var actualCost float64
				actualCost += float64(cartItem.Quantity) * fruit.Price
				discount := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(singleItemDiscount.Discount)
				calculatedCost = (actualCost - discount)
				appliedSingleItemDiscount.Savings = discount
				models.DB.Model(&cartItem).Update("item_total", calculatedCost).Update("item_discounted_total", discount)
				if err := models.DB.Model(&appliedSingleItemDiscount).
					Where("cart_id = ?", cartItem.CartID).
					First(&appliedSingleItemDiscount).Error; err != nil {
					if gorm.IsRecordNotFoundError(err) {
						models.DB.Create(&appliedSingleItemDiscount)
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

}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyDualItemDiscount(cartItem models.CartItem) {
	//, "fruit_id_1 in (?) or fruit_id_2 in (?)", cartItem.FruitID, cartItem.FruitID
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount []models.DualItemDiscount
	models.DB.Where("fruit_id_1 = ? or fruit_id_2 = ?", cartItem.FruitID, cartItem.FruitID).Find(&dualItemDiscount)

	x := make([]models.DualItemDiscount, 0)
	for _, dualItemDiscount := range dualItemDiscount {
		x = append(x, dualItemDiscount)
		appliedDualItemDiscount := models.AppliedDualItemDiscount{
			CartID: cartItem.CartID,
			//DualItemDiscount: dualItemDiscount,
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
				if cartItem.FruitID == dualItemDiscount.FruitID {
					var fruit models.Fruit
					models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
					discount := float64(sets*dualItemDiscount.Count_1) / float64(100) * float64(dualItemDiscount.Discount)
					appliedDualItemDiscount.Savings += discount
					models.DB.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
						Update("item_discounted_total", discount)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					var fruit models.Fruit
					models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
					discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
					appliedDualItemDiscount.Savings += discount
					models.DB.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
						Update("item_discounted_total", discount)
				}
			}
			discountUpdate = true
		} else {
			for _, x := range cartItems {
				cartItem := x
				if cartItem.FruitID == dualItemDiscount.FruitID {
					models.DB.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("item_discounted_total", 0.0)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					models.DB.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("tem_discounted_total", 0.0)
				}
			}

		}

		if discountUpdate {
			appliedDualItemDiscount.DualItemDiscount = x
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
}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
/* func ApplyDualItemDiscount(cartItem models.CartItem, discountCouponCode string) {
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount []models.DualItemDiscount
	models.DB.Where("name = ?", discountCouponCode).Find(&dualItemDiscount)

	appliedDualItemDiscount := models.AppliedDualItemDiscount{
		CartID:           cartItem.CartID,
		DualItemDiscount: dualItemDiscount,
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

} */

// getSets gets the number of sets of banana and pears
func getSets(fruit1 int, fruit2 int, count1 int, count2 int) int {
	fmt.Println("fruit1 ", fruit1)
	fmt.Println("fruit2 ", fruit2)
	fmt.Println("count1 ", count1)
	fmt.Println("count2 ", count2)

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
