package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// CreateUpdateItemInCart will either create an item in the cart or remove based the quantity provided in the payload
func (server *Server) CreateUpdateItemInCart(w http.ResponseWriter, r *http.Request) {
	// Reading the request body from http request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// Creating cartItem struct mapped from the request payloads
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
	// Save the cart item entry into the database
	createdCartItem, err := cartItem.SaveOrUpdateCartItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	ApplySingleItemDiscounts(*createdCartItem, server.DB)
	ApplyDualItemDiscounts(cartItem, server.DB)
	RecalcualtePayments(server.DB, cartItem.CartID)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCartItem.ID))
	responses.JSON(w, http.StatusCreated, createdCartItem)
}

// GetCartItems fetched all items in a given cart id
func (server *Server) GetCartItems(w http.ResponseWriter, r *http.Request) {

	cartItem := models.CartItem{}
	vars := mux.Vars(r)
	cart_id := vars["cart_id"]

	cartItems := cartItem.FindAllCartItems(server.DB, cart_id)

	responses.JSON(w, http.StatusOK, cartItems)
}

// RecalcualtePayments recalcuates the cart value and its saving based on the cart items
func RecalcualtePayments(db *gorm.DB, cartID uint) {
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
}

// ApplySingleItemDiscounts applies all single item discounts based on the discounts added to the table(single_item_discounts),
// for this instance there is only on coupon which APPLE10. This feature is implemented in this way so that in future if we ever
// wanted to add new single item discounts we can update the single_item_discounts table and no need to update the business logic
// around how to calculate the discount as the single_item_discounts table carries the business rule information such as
// quantity, discount percentage and fruit to the discount needs to be applied.
func ApplySingleItemDiscounts(cartItem models.CartItem, db *gorm.DB) {
	// Preloading the available discounts of the given fruit
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
			if cartItem.Quantity >= singleItemDiscount.Count {
				var costAfterDiscount float64
				var actualCost float64
				actualCost += float64(cartItem.Quantity) * fruit.Price
				discount := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(singleItemDiscount.Discount)
				costAfterDiscount = (actualCost - discount)
				appliedSingleItemDiscount.Savings = discount
				db.Model(&cartItem).Update("item_total", costAfterDiscount).Update("item_discounted_total", discount)
				if err := db.Model(&appliedSingleItemDiscount).
					Where("cart_id = ?", cartItem.CartID).
					First(&appliedSingleItemDiscount).Error; err != nil {
					if gorm.IsRecordNotFoundError(err) {
						db.Create(&appliedSingleItemDiscount)
					}
				}
			} else {
				db.Unscoped().Where("cart_id = ?", cartItem.CartID).Delete(&appliedSingleItemDiscount)
			}
		}
	}
}

// ApplyDualItemDiscounts applies all dual item discounts based on the discounts added to the table(dual_item_discounts),
// for this instance there is only on coupon which PEARBANANA30. This feature is implemented in this way so that in future if we ever
// wanted to add new double item discounts we can update the dual_item_discounts table and no need to update the business logic
// around how to calculate the discount as the dual_item_discounts table carries the business rule information such as
// quantity, discount percentage and fruits  to the discount needs to be applied.
func ApplyDualItemDiscounts(cartItem models.CartItem, db *gorm.DB) {
	var fruit models.Fruit
	db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount []models.DualItemDiscount
	db.Where("fruit_id_1 = ? or fruit_id_2 = ?", cartItem.FruitID, cartItem.FruitID).Find(&dualItemDiscount)

	x := make([]models.DualItemDiscount, 0)
	for _, dualItemDiscount := range dualItemDiscount {
		x = append(x, dualItemDiscount)
		appliedDualItemDiscount := models.AppliedDualItemDiscount{
			CartID: cartItem.CartID,
		}
		var fruitId1Count int
		var fruitId2Count int
		var discountUpdate bool

		var cartItems []models.CartItem
		if err := db.Where("cart_id = ?", cartItem.CartID).Find(&cartItems).Error; err != nil {
			fmt.Println("Error ", err)
		}
		// Count number of fruits per each type in dual discount category
		for _, x := range cartItems {
			cartItem := x
			if cartItem.FruitID == dualItemDiscount.FruitID_1 {
				fruitId1Count = cartItem.Quantity

			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				fruitId2Count = cartItem.Quantity
			}
		}

		if fruitId1Count > 0 && fruitId2Count > 0 {
			// get the sets formation of given  discounted fruits available
			sets := getSets(fruitId1Count, fruitId2Count, dualItemDiscount.Count_1, dualItemDiscount.Count_2)
			if sets != 0 {
				for _, x := range cartItems {
					cartItem := x
					if cartItem.FruitID == dualItemDiscount.FruitID_1 {
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

// getSets gets the number of sets of given two fruits using given reference counts
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
	return set
}
