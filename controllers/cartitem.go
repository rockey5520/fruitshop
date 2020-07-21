package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CartItemInput struct {
	CartId uint   `json:"cartid" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Count  int    `json:"count"`
}

// @Summary Show details of a all items in a cart
// @Description Get details of contents of the cart
// @Accept  json
// @Produce  json
// @Param cart_id path string true "Customer identifier"
// @Success 200 {object} models.CartItem
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/cartitem/{cart_id} [get]
// GetAllCartList will retuen all of the cart list for a given customer
func GetAllCartItems(c *gin.Context) {
	s := make([]models.CartItemResponse, 0)

	var cartItems []models.CartItem
	models.DB.Where("cart_id = ?", c.Param("cart_id")).Find(&cartItems)

	for _, cartItem := range cartItems {
		var fruit models.Fruit
		models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
		s = append(s, models.CartItemResponse{
			Name:        fruit.Name,
			CostPerItem: fruit.Price,
			Count:       cartItem.Quantity,
			ItemTotal:   cartItem.ItemTotal,
		})

	}
	c.JSON(http.StatusOK, gin.H{"data": s})
}

// @Summary Creates/Updated item in the cart
// @Description This end point will record cart item details into the database
// @Accept  json
// @Produce  json
// @Param Input body models.CartItem true "Input request"
// @Param login_id path string true "Customer identifier"
// @Success 200 {object} models.CartItem
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/cartitem/{login_id} [post]
// CreateUpdateItemInCart will add users choosen fruits to the cart list
func CreateUpdateItemInCart(c *gin.Context) {
	// Bind the input payload to schema for validations
	var input CartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fruit models.Fruit
	if err := models.DB.Where("name = ?", input.Name).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: input.CartId, FruitID: fruit.ID, ItemTotal: fruit.Price * float64(input.Count)}
	if input.Count > 0 {
		// Create/update fruit to the cart
		cartItem.Quantity = input.Count
		if err := models.DB.Model(&cartItem).Where("cart_id = ? and fruit_id = ?", input.CartId, fruit.ID).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&cartItem) // create new record from newUser
			}
		} else {
			models.DB.Model(&cartItem).Where("cart_id = ? and fruit_id = ?", input.CartId, fruit.ID).Update("quantity", input.Count).Update("item_total", float64(input.Count)*fruit.Price)
		}
	} else if input.Count == 0 {
		models.DB.Where("cart_id = ? and fruit_id = ?", input.CartId, fruit.ID).Delete(&cartItem)

	}

	// RecalcuateItem payment for the item in the cart
	ApplySingleItemDiscounts(cartItem)
	ApplyDualItemDiscounts(cartItem)
	// Recalcuate the payment for the cart
	RecalcualtePayments(cartItem.CartID)

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}

// ApplySingleItemDiscounts applies single fruit discounts based on the single item discounts
func ApplySingleItemDiscounts(cartItem models.CartItem) {

	ApplyApple10Discount(cartItem)

}

// ApplyDualItemDiscounts applies single fruit discounts based on the dual item combo  discounts
func ApplyDualItemDiscounts(cartItem models.CartItem) {

	//Applying pear banana 30% discount
	ApplyBananaPear30Discount(cartItem)

}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(cartID uint) {
	// Recalcualte the payments

	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
	}
	var cart models.Cart
	if err := models.DB.Where("ID = ?", cartID).Find(&cart).Error; err != nil {
		fmt.Println("Error ", err)
	}
	models.DB.Model(&cart).Update("total", totalCost)
	var payment models.Payment
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&payment).Error; err != nil {
		fmt.Println("Error ", err)
	}
	models.DB.Model(&payment).Update("amount", totalCost)

}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyBananaPear30Discount(cartItem models.CartItem) {
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var dualItemDiscount models.DualItemDiscount
	models.DB.Where("name = ?", "PEARBANANA30").Find(&dualItemDiscount)

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
		fmt.Println("cartItem.FruitID", cartItem.FruitID)
		fmt.Println("dualItemDiscount.FruitID_1", dualItemDiscount.FruitID_1)
		fmt.Println("dualItemDiscount.FruitID_2", dualItemDiscount.FruitID_2)
		if cartItem.FruitID == dualItemDiscount.FruitID_1 {
			fruitId1Count = cartItem.Quantity

		} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
			fruitId2Count = cartItem.Quantity
		}
	}
	fmt.Println("fruitId1Count ", fruitId1Count)
	fmt.Println("fruitId2Count ", fruitId2Count)

	sets := getSets(fruitId1Count, fruitId2Count, dualItemDiscount.Count_1, dualItemDiscount.Count_2)
	fmt.Println("sets", sets)
	if sets != 0 {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.FruitID == dualItemDiscount.FruitID_1 {
				var fruit models.Fruit
				models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_1) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings = discount
				models.DB.Model(&cartItem).Where("cart_id = ?", cartItem.CartID).Update("item_total", (float64(cartItem.Quantity)*fruit.Price)-discount)
			} else if cartItem.FruitID == dualItemDiscount.FruitID_2 {
				var fruit models.Fruit
				models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
				discount := float64(sets*dualItemDiscount.Count_2) / float64(100) * float64(dualItemDiscount.Discount)
				appliedDualItemDiscount.Savings = discount
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

// getSets gets the number of sets of banana and pears
func getSets(fruit1 int, fruit2 int, count1 int, count2 int) int {

	fruit1Count := fruit1
	fruit2Count := fruit2
	var set int
	fmt.Println("fruit1Count", fruit1Count)
	fmt.Println("fruit2Count", fruit2Count)
	fmt.Println("count1", count1)
	fmt.Println("count2", count2)
	fmt.Println("====================")

	for i := 0; i < fruit1Count; i++ {
		fmt.Println("i", i)
		fmt.Println("fruit1Count", fruit1Count)
		fmt.Println("fruit2Count", fruit2Count)
		fmt.Println("count1", count1)
		fmt.Println("count2", count2)
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

//ApplyApple10Discount applies apple 10 percent discount
func ApplyApple10Discount(cartItem models.CartItem) {
	var fruit models.Fruit
	models.DB.Where("ID = ?", cartItem.FruitID).Find(&fruit)
	var singeItemDiscount models.SingleItemDiscount
	models.DB.Where("fruit_id = ?", fruit.ID).Find(&singeItemDiscount)
	appliedSingleItemDiscount := models.AppliedSingleItemDiscount{
		CartID:               cartItem.CartID,
		SingleItemDiscountID: singeItemDiscount.ID,
	}
	if cartItem.FruitID == fruit.ID && singeItemDiscount.Name == "APPLE10" {
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
