package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"

	"github.com/jinzhu/gorm"
)

// CreateCustomer is
func (server *Server) CreateUpdateItemInCart(w http.ResponseWriter, r *http.Request) {
	// Reading the request body from http request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// Creating customer , cart structs and mapping request body to customer and a creating new card with customer ID
	cartItem := models.CartItem{}
	err = json.Unmarshal(body, &cartItem)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Customer validation
	err = cartItem.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	createdCartItem, err := cartItem.SaveOrUpdateCartItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	ApplySingleItemDiscounts(*createdCartItem, server.DB)
	ApplyDualItemDiscounts(cartItem, server.DB)
	RecalcualtePayments(cartItem.CartID, server.DB)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCartItem.ID))
	responses.JSON(w, http.StatusCreated, createdCartItem)
}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(cartID uint, db *gorm.DB) {
	// Recalcualte the payments
	var cartItems []models.CartItem
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	var totalDiscountedCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
		totalDiscountedCost += item.ItemDiscountedTotal
	}
	var cart models.Cart
	if err := db.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}
	db.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
	/* var payment models.Payment
	if err := db.Where("cart_id = ?", cart.ID).Find(&payment).Error; err != nil {
		fmt.Println("Error ", err)
	}
	db.Model(&payment).Where("cart_id = ?", cart.ID).Update("amount", totalCost) */

}

//ApplyApple10Discount applies apple 10 percent discount
func ApplySingleItemDiscounts(cartItem models.CartItem, db *gorm.DB) {
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
					Where("cart_id = ?", cartItem.CartID).
					First(&appliedSingleItemDiscount).Error; err != nil {
					if gorm.IsRecordNotFoundError(err) {
						db.Create(&appliedSingleItemDiscount)
					}
				}
			} else {
				var itemCost float64
				itemCost += float64(cartItem.Quantity) * fruit.Price
				db.Model(&cartItem).Update("item_total", itemCost)
				db.Unscoped().Where("cart_id = ?", cartItem.CartID).Delete(&appliedSingleItemDiscount)
			}
		}
	}

}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyDualItemDiscounts(cartItem models.CartItem, db *gorm.DB) {
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
		if err := db.Where("cart_id = ?", cartItem.CartID).Find(&cartItems).Error; err != nil {
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
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
						Update("item_discounted_total", discount)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					var fruit models.Fruit
					db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
					discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
					appliedDualItemDiscount.Savings += discount
					db.Model(&cartItem).
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
					db.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("item_discounted_total", 0.0)
				} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
					db.Model(&cartItem).
						Where("cart_id = ?", cartItem.CartID).
						Update("item_total", (float64(cartItem.Quantity)*fruit.Price)).
						Update("tem_discounted_total", 0.0)
				}
			}

		}

		if discountUpdate {
			appliedDualItemDiscount.DualItemDiscount = x
			if err := db.Model(&appliedDualItemDiscount).
				Where("cart_id = ?", cartItem.CartID).
				First(&appliedDualItemDiscount).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					db.Create(&appliedDualItemDiscount) // create new record from newUser
				}
			}
		} else {
			db.Unscoped().Where("cart_id = ?", cartItem.CartID).Delete(&appliedDualItemDiscount)
		}
	}
}

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
