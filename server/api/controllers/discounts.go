package controllers

import (
	"fmt"
	"fruitshop/server/api/models"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// @Summary Show details of all discounts for a given cart
// @Description Get details of all discounts for each cart
// @Accept  json
// @Produce  json
// @Param cartID path string true "Customer identifier"
// @Success 200 {array} models.Discount
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/discounts/{cartID} [get]

// FindDiscounts will return all discounts with status as APPLIED available within the fruitshop
func (server *Server) FindDiscounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, err := vars["cart_id"]
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	db.Where("ID = ?", cartID).Find(&cart)
	appliedDiscountsResponseList := make([]models.AppliedDiscountsResponse, 0)

	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{}
	db.Where("cartID = ?", cartID).
		Preload("SingleItemDiscount").
		Find(&appliedSingleItemDiscount)
	appliedDualItemDiscount := models.AppliedDualItemDiscount{}
	db.Where("cartID = ?", cartID).
		Preload("DualItemDiscount").
		Find(&appliedDualItemDiscount)
	appliedSingleItemCoupon := models.AppliedSingleItemCoupon{}
	db.Where("cartID = ?", cart.ID).
		Find(&appliedSingleItemCoupon)
	var singeItemCoupon models.SingleItemCoupon
	db.Where("ID = ?", appliedSingleItemCoupon.SingleItemCouponID).Find(&singeItemCoupon)

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

	/* 	if singeItemCoupon.Name != "" {
		s = append(s, models.Discount{
			Name:   singeItemCoupon.Name,
			Status: "APPLIED",
		})
	} */
	//for _, singeItemCoupon := range appliedSingleItemCoupon.SingleItemCoupon {
	if singeItemCoupon.Name != "" {
		fmt.Println("singeItemCoupon.Name", singeItemCoupon.Name)
		appliedDiscountsResponseList = append(appliedDiscountsResponseList, models.AppliedDiscountsResponse{
			Name:   singeItemCoupon.Name,
			Status: "APPLIED",
		})
	}
	//}

	responses.JSON(w, http.StatusOK, appliedDiscountsResponseList)
}

// @Summary Applied orange 30 coupon code
// @Description This endpoint applied orange 30 percent discount coupon code
// @Accept  json
// @Produce  json
// @Param cartID path string true "Customer identifier"

// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/orangecoupon/{cartID} [get]
//ApplyOrangeCoupon applies discount for oranges
func ApplyTimeSensitiveCoupon(c *gin.Context) {
	// This coupon will be run as go routine and sleeps for 10 seconds
	go ApplySingleItemTimSensitiveCoupon(c)
}

func ApplySingleItemTimSensitiveCoupon(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID, _ := strconv.Atoi(cartID)
	fruit_id, _ := strconv.Atoi(c.Param("fruit_id"))

	fruit := models.Fruit{}
	db.Where("ID = ?", fruit_id).
		Preload("SingleItemCoupon").
		Find(&fruit)

	//singleItemCouponList := fruit.SingleItemCoupon
	var singleItemCouponList []models.SingleItemCoupon
	db.Where("fruit_id = ?", fruit_id).Find(&singleItemCouponList)
	var interestCoupon models.SingleItemCoupon
	for _, coupon := range singleItemCouponList {
		if coupon.FruitID == fruit.ID {
			interestCoupon = coupon
		}
	}
	fmt.Println("cartID", cartID)
	fmt.Println("uint cartID", uint(cartID))
	appliedSingleItemCoupon := models.AppliedSingleItemCoupon{
		CartID:             uint(cartID),
		SingleItemCouponID: interestCoupon.ID,
		//SingleItemCoupon: singleItemCouponList,
	}
	if err := db.Model(&appliedSingleItemCoupon).
		Where("cartID = ?", cartID).
		First(&appliedSingleItemCoupon).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.Create(&appliedSingleItemCoupon)
		}
	} else {
		db.Model(&appliedSingleItemCoupon).
			Where("cartID = ? ", cartID).
			Update("single_item_coupon_id", interestCoupon.ID)
	}

	var cartItem models.CartItem
	db.Where("cartID = ? AND fruit_id = ?", cartID, fruit_id).Find(&cartItem)

	/* var fruit models.Fruit
	db.Where("ID = ?", cartItem.FruitID).Find(&fruit) */
	if cartItem.Quantity > 0 {
		discountCalculated := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(interestCoupon.Discount)
		updatedTotalCost := cartItem.ItemTotal - discountCalculated
		db.Model(&cartItem).
			Where("cartID = ?", cartItem.CartID).
			Update("ItemTotal", updatedTotalCost).
			Update("item_discounted_total", discountCalculated)
		RecalcualtePayments(cartItem.CartID, c)
		appliedSingleItemCoupon.Savings = discountCalculated
		if err := db.Model(&appliedSingleItemCoupon).
			Where("cartID = ?", cartItem.CartID).
			First(&appliedSingleItemCoupon).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&appliedSingleItemCoupon)
			}
		} else {
			db.Model(&appliedSingleItemCoupon).
				Where("cartID = ? ", cartID).
				Update("savings", discountCalculated)
		}
	}

	// configurable timer for the coupon expiry
	time.Sleep(10 * time.Second)

	var cart models.Cart
	db.Where("ID = ?", cartID).Find(&cart)
	if cart.Status != "CLOSED" {
		var cartItem models.CartItem
		db.Where("ID = ?", cartID).First((&cartItem))
		var fruit models.Fruit
		db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		db.Model(&cartItem).
			Where("cartID = ?", cartItem.ID).
			Update("ItemTotal", float64(cartItem.Quantity)*fruit.Price).
			Update("item_discounted_total", 0.0)
		RecalcualtePayments(cartItem.CartID, c)
		db.Unscoped().Where("cartID = ?", cartItem.CartID).Delete(&appliedSingleItemCoupon)
	}

}

// ApplySingleItemDiscounts applies single fruit discounts based on the single item discounts
func ApplySingleItemDiscounts(cartItem models.CartItem, c *gin.Context) {

	// This applies 10 percent discount for APPLE and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplySingleItemDiscount(cartItem, c)

}

// ApplyDualItemDiscounts applies discount for dual items discounts based on the dual item combo  discounts
func ApplyDualItemDiscounts(cartItem models.CartItem, c *gin.Context) {

	// This applies 30% discount for Banana pear set and the reason i encapsulated is so in future without change in business logic
	// change to calucate discounts on single unique item discounts so our application can be pushed to production
	// sooner than to build a new logic for each discount coupon code
	ApplyDualItemDiscount(cartItem, c)

}

//ApplyApple10Discount applies apple 10 percent discount
func ApplySingleItemDiscount(cartItem models.CartItem, c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	fruit := models.Fruit{}
	db.Where("ID = ?", cartItem.FruitID).
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
				db.Model(&cartItem).Update("item_total", calculatedCost).Update("item_discounted_total", discount)
				if err := db.Model(&appliedSingleItemDiscount).
					Where("cartID = ?", cartItem.CartID).
					First(&appliedSingleItemDiscount).Error; err != nil {
					if gorm.IsRecordNotFoundError(err) {
						db.Create(&appliedSingleItemDiscount)
					}
				}
			} else {
				var itemCost float64
				itemCost += float64(cartItem.Quantity) * fruit.Price
				db.Model(&cartItem).Update("item_total", itemCost)
				db.Unscoped().Where("cartID = ?", cartItem.CartID).Delete(&appliedSingleItemDiscount)
			}
		}
	}

}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyDualItemDiscount(cartItem models.CartItem, c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//, "fruit_id_1 in (?) or fruit_id_2 in (?)", cartItem.FruitID, cartItem.FruitID
	var fruit models.Fruit
	db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount []models.DualItemDiscount
	db.Where("fruit_id_1 = ? or fruit_id_2 = ?", cartItem.FruitID, cartItem.FruitID).Find(&dualItemDiscount)

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
		if err := db.Where("cartID = ?", cartItem.CartID).Find(&cartItems).Error; err != nil {
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
					db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
					discount := float64(sets*dualItemDiscount.Count_1) / float64(100) * float64(dualItemDiscount.Discount)
					appliedDualItemDiscount.Savings += discount
					db.Model(&cartItem).
						Where("cartID = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
						Update("item_discounted_total", discount)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					var fruit models.Fruit
					db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
					discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
					appliedDualItemDiscount.Savings += discount
					db.Model(&cartItem).
						Where("cartID = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
						Update("item_discounted_total", discount)
				}
			}
			discountUpdate = true
		} else {
			for _, x := range cartItems {
				cartItem := x
				if cartItem.FruitID == dualItemDiscount.FruitID {
					db.Model(&cartItem).
						Where("cartID = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("item_discounted_total", 0.0)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					db.Model(&cartItem).
						Where("cartID = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("tem_discounted_total", 0.0)
				}
			}

		}

		if discountUpdate {
			appliedDualItemDiscount.DualItemDiscount = x
			if err := db.Model(&appliedDualItemDiscount).
				Where("cartID = ?", cartItem.CartID).
				First(&appliedDualItemDiscount).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					db.Create(&appliedDualItemDiscount) // create new record from newUser
				}
			}
		} else {
			db.Unscoped().Where("cartID = ?", cartItem.CartID).Delete(&appliedDualItemDiscount)
		}
	}
}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
/* func ApplyDualItemDiscount(cartItem models.CartItem, discountCouponCode string) {
	var fruit models.Fruit
	db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount []models.DualItemDiscount
	db.Where("name = ?", discountCouponCode).Find(&dualItemDiscount)

	appliedDualItemDiscount := models.AppliedDualItemDiscount{
		CartID:           cartItem.CartID,
		DualItemDiscount: dualItemDiscount,
	}

	var fruitId1Count int
	var fruitId2Count int
	var discountUpdate bool

	var cartItems []models.CartItem
	if err := db.Where("cartID = ?", cartItem.CartID).Find(&cartItems).Error; err != nil {
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
				db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_1) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings += discount
				db.Model(&cartItem).Where("cartID = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount)
			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				var fruit models.Fruit
				db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings += discount
				db.Model(&cartItem).Where("cartID = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount)
			}
		}
		discountUpdate = true
	} else {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.FruitID == dualItemDiscount.FruitID_1 {
				db.Model(&cartItem).Where("cartID = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity) * fruit.Price))
			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				db.Model(&cartItem).Where("cartID = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity) * fruit.Price))
			}
		}

	}

	if discountUpdate {
		if err := db.Model(&appliedDualItemDiscount).
			Where("cartID = ?", cartItem.CartID).
			First(&appliedDualItemDiscount).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				db.Create(&appliedDualItemDiscount) // create new record from newUser
			}
		}
	} else {
		db.Unscoped().Where("cartID = ?", cartItem.CartID).Delete(&appliedDualItemDiscount)
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
