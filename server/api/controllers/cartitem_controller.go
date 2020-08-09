package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"fruitshop/api/models"
	"fruitshop/api/responses"
	"fruitshop/api/utils/formaterror"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// CreateItemInCart will create an item in the cart based the quantity provided in the payload
func (server *Server) CreateItemInCart(w http.ResponseWriter, r *http.Request) {
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
	createdCartItem, err := cartItem.SaveCartItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	// Applying business rules( such as single item discounts for Apple, and dual item discounts for Pear and Banana)
	ApplySingleItemDiscounts(*createdCartItem, server.DB)
	ApplyDualItemDiscounts(cartItem, server.DB)

	// Recalculate Cart value , Discounts post adding/deleting item from cart and applying business rules
	RecalcualtePayments(server.DB, cartItem.CartID)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCartItem.ID))
	responses.JSON(w, http.StatusCreated, createdCartItem)
}

// UpdateItemInCart will update  item in the cart based the quantity provided in the payload
func (server *Server) UpdateItemInCart(w http.ResponseWriter, r *http.Request) {
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
	createdCartItem, err := cartItem.UpdateCartItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	// Applying business rules( such as single item discounts for Apple, and dual item discounts for Pear and Banana)
	ApplySingleItemDiscounts(*createdCartItem, server.DB)
	ApplyDualItemDiscounts(cartItem, server.DB)

	// Recalculate Cart value , Discounts post adding/deleting item from cart and applying business rules
	RecalcualtePayments(server.DB, cartItem.CartID)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCartItem.ID))
	responses.JSON(w, http.StatusCreated, createdCartItem)
}

// DeleteItemInCart will remove based the quantity provided in the payload
func (server *Server) DeleteItemInCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cart_id := vars["cart_id"]
	fruitname := vars["fruitname"]
	fmt.Println("fruit_name ", fruitname)
	cartID, _ := strconv.Atoi(cart_id)
	//fruitID, _ := strconv.Atoi(fruit_id)
	cartItem := models.CartItem{
		CartID: uint(cartID),
		//FruitID:             uint(fruitID),
		Name:                fruitname,
		Quantity:            0,
		ItemTotal:           0,
		ItemDiscountedTotal: 0,
	}

	// Save the cart item entry into the database
	createdCartItem, err := cartItem.DeleteCartItem(server.DB, cart_id, fruitname)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	// Applying business rules( such as single item discounts for Apple, and dual item discounts for Pear and Banana)
	ApplySingleItemDiscounts(*createdCartItem, server.DB)
	ApplyDualItemDiscounts(cartItem, server.DB)

	// Recalculate Cart value , Discounts post adding/deleting item from cart and applying business rules
	RecalcualtePayments(server.DB, cartItem.CartID)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, createdCartItem.ID))
	responses.JSON(w, http.StatusCreated, createdCartItem)
}

// GetCartItems fetches all items in a given cart id
func (server *Server) GetCartItems(w http.ResponseWriter, r *http.Request) {
	// Reading cart_id from request params
	vars := mux.Vars(r)
	cart_id := vars["cart_id"]
	cartItem := models.CartItem{}
	cartItems := cartItem.FindAllCartItems(server.DB, cart_id)
	responses.JSON(w, http.StatusOK, cartItems)
}

// RecalcualtePayments recalcuates the cart value and its saving based on the cart items
func RecalcualtePayments(db *gorm.DB, cartID uint) {
	// Fetch all items in a given cart
	var cartItems []models.CartItem
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}

	// calcuate the total cost of the cartitems and total discounts applied
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

	// Update Cart table with total cost and total savings
	db.Model(&cart).Update("total", totalCost).Update("total_savings", totalDiscountedCost)
}

/*
   ApplySingleItemDiscounts applies all single item discounts based on the discounts added to the meta table(single_item_discounts),
   at this point of time there is only on coupon which APPLE10. I have implemented this feature in this way so that in future if we ever
   wanted to add new single item discounts we can update the single_item_discounts table and no need to update the business logic
   around how to calculate the discount as the single_item_discounts table carries the business rule information such as
   quantity, discount percentage and fruit to the discount needs to be applied and this function works like magic
*/
func ApplySingleItemDiscounts(cartItem models.CartItem, db *gorm.DB) {
	// Preloading the available discounts for a given fruit
	fruit := models.Fruit{}
	db.Where("ID = ?", cartItem.FruitID).
		Preload("SingleItemDiscount").
		Find(&fruit)

	singeItemDiscount := fruit.SingleItemDiscount
	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{
		CartID: cartItem.CartID,
	}

	/*
		For each discount rule loaded to the table we will check if the given cartitem can satisfy the conditions stated
		and updated the appliedSingleItemDiscount table with the information that discount is applied and also updates the
		CartItem table with updated cost and savings calculated due the discount rules. When an Item is ammended in cart
		this function again removes entry in appliedSingleItemDiscount table and corrects the total cost according to the
		price set in the fruits table

	*/
	for _, singleItemDiscount := range singeItemDiscount {
		if cartItem.FruitID == singleItemDiscount.FruitID {
			if cartItem.Quantity >= singleItemDiscount.Count {
				var costAfterDiscount float64
				var actualCost float64
				actualCost += float64(cartItem.Quantity) * fruit.Price
				discount := ((float64(cartItem.Quantity) * fruit.Price) / 100) * float64(singleItemDiscount.Discount)
				costAfterDiscount = (actualCost - discount)
				appliedSingleItemDiscount.Savings = discount
				appliedSingleItemDiscount.SingleItemDiscountID = singleItemDiscount.ID
				appliedSingleItemDiscount.SingleItemDiscountName = singleItemDiscount.Name
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

/*
    For each discount rule loaded to the table we will check if the given cartitem can satisfy the conditions stated
	and updated the appliedDualItemDiscount table with the information that discount is applied and also updates the
	CartItem table with updated cost and savings calculated due the discount rules. When an Item is ammended in cart
	this function again removes entry in appliedDualItemDiscount table and corrects the total cost according to the
	price set in the fruits table
*/
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
						appliedDualItemDiscount.DualItemDiscountID = dualItemDiscount.ID
						appliedDualItemDiscount.DualItemDiscountName = dualItemDiscount.Name
						db.Model(&cartItem).
							Where("cart_id = ?", cartItem.CartID).
							Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount).
							Update("item_discounted_total", discount)
					} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
						var fruit models.Fruit
						db.Where("ID = ?", cartItem.FruitID).Find(&fruit)
						discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
						appliedDualItemDiscount.Savings += discount
						appliedDualItemDiscount.DualItemDiscountID = dualItemDiscount.ID
						appliedDualItemDiscount.DualItemDiscountName = dualItemDiscount.Name
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
