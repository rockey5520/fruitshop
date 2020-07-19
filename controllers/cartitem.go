package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateUpdateItemInCart will add users choosen fruits to the cart list
func CreateUpdateItemInCart(c *gin.Context) {
	// Bind the input payload to schema for validations
	var input CartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Fetch customer information
	var customer models.Customer
	if err := models.DB.Where("login_id = ?", c.Param("login_id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer record not found!"})
		return
	}
	// Fetch cart information
	var cart models.Cart
	if err := models.DB.Where("customer_id = ?", customer.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart record not found!"})
		return
	}
	var fruit models.Fruit
	if err := models.DB.Where("name = ?", input.Name).First(&fruit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fruit record not found!"})
		return
	}
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: cart.ID, Name: input.Name, CostPerItem: fruit.Price}
	if input.Count > 0 {
		// Create/update fruit to the cart

		cartItem.Count = input.Count
		if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				models.DB.Create(&cartItem) // create new record from newUser
			}
		} else {
			models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).Update("count", input.Count)
		}
	} else if input.Count == 0 {
		models.DB.Where("cart_id = ? AND name = ?", cart.ID, input.Name).Delete(&cartItem)
		//if err := models.DB.Model(&cartItem).Where("cart_id = ? and name = ?", cart.ID, input.Name).First(&cartItem).Error; err != nil {

	}
	models.DB.Where("name = ?", cartItem.Name).Find(&fruit)
	// RecalcuateItem payment for the item in the cart
	RecalcualteIemPayments(customer, cart, cartItem, fruit)
	// Recalcuate the payment for the cart
	RecalcualtePayments(cart)

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}

// RecalcualteItemPayments recalcuates the payment for the cart
func RecalcualteIemPayments(customer models.Customer, cart models.Cart, cartItem models.CartItem, fruit models.Fruit) {
	// Recalcualte the payments
	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var discounts []models.Discount
	if err := models.DB.Where("customer_id = ?", customer.ID).Find(&discounts).Error; err != nil {
		fmt.Println("Error ", err)
	}

	for _, item := range cartItems {
		fmt.Println(item.Name)
		if item.Name == "Apple" {
			if cartItem.Count >= 7 {
				ApplyApple10Discount(customer, cart, cartItem, fruit)
			} else {
				var discount models.Discount
				models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "APPLE10").Update("status", "NOTAPPLIED")
				models.DB.Model(&cartItem).Where("cart_id = ? AND name = ?", cart.ID, "Apple").Update("item_total", (float64(cartItem.Count) * fruit.Price))
			}
		} else if item.Name == "Banana" || item.Name == "Pear" {
			//Applying pear banana 30% discount
			ApplyBananaPear30Discount(customer, cart, cartItem, fruit)
		} else if item.Name == "Orange" {
			models.DB.Model(&cartItem).Where("cart_id = ? AND name = ?", cart.ID, "Orange").Update("item_total", (float64(cartItem.Count) * fruit.Price))
		}
	}

}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(cart models.Cart) {
	// Recalcualte the payments

	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	for _, item := range cartItems {
		totalCost += item.ItemTotal
	}
	models.DB.Model(&cart).Update("total", totalCost)
}

//ApplyBananaPear30Discount applies banana pear 30 percent discount
func ApplyBananaPear30Discount(customer models.Customer, cart models.Cart, cartItem models.CartItem, fruit models.Fruit) {
	var pearCount int
	var bananaCount int
	var discountUpdate bool

	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	for _, x := range cartItems {
		cartItem := x
		if cartItem.Name == "Pear" {
			pearCount = cartItem.Count

		} else if cartItem.Name == "Banana" {
			bananaCount = cartItem.Count
		}
	}
	sets := getSets(pearCount, bananaCount)
	if sets != 0 {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.Name == "Pear" {
				var fruit models.Fruit
				models.DB.Where("name = ?", cartItem.Name).Find(&fruit)
				discount := float64(sets*4) / float64(100) * float64(30)
				models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", (float64(cartItem.Count)*fruit.Price)-discount)
			} else if cartItem.Name == "Banana" {
				var fruit models.Fruit
				models.DB.Where("name = ?", cartItem.Name).Find(&fruit)
				discount := float64(sets*2) / float64(100) * float64(30)
				models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", (float64(cartItem.Count)*fruit.Price)-discount)
			}
		}
		discountUpdate = true
	} else {
		for _, x := range cartItems {
			cartItem := x
			if cartItem.Name == "Pear" {
				models.DB.Model(&cartItem).Where("cart_id = ? AND name = ?", cart.ID, "Pear").Update("item_total", (float64(cartItem.Count) * fruit.Price))
			} else if cartItem.Name == "Banana" {
				models.DB.Model(&cartItem).Where("cart_id = ? AND name = ?", cart.ID, "Banana").Update("item_total", (float64(cartItem.Count) * fruit.Price))
			}
		}

	}

	if discountUpdate {
		var discount models.Discount
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "PEARBANANA").Update("status", "APPLIED")
	} else {
		var discount models.Discount
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "PEARBANANA").Update("status", "NOTAPPLIED")
	}

}

//ApplyApple10Discount applies apple 10 percent discount
func ApplyApple10Discount(customer models.Customer, cart models.Cart, cartItem models.CartItem, fruit models.Fruit) {
	var itemCost float64
	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var appleCost float64
	appleCost += float64(cartItem.Count) * fruit.Price

	discount := ((float64(cartItem.Count) * fruit.Price) / 100) * 10
	itemCost = (appleCost - discount)

	models.DB.Model(&cartItem).Where("cart_id = ?", cart.ID).Update("item_total", itemCost)
	var discountItem models.Discount
	models.DB.Model(&discountItem).Where("customer_id = ? AND name = ?", cart.CustomerId, "APPLE10").Update("status", "APPLIED")

}

// getSets gets the number of sets of banana and pears
func getSets(pear int, banana int) int {

	pearCount := pear
	bananaCount := banana
	var set int

	for i := 0; i < pearCount; i++ {
		if pearCount >= 4 && bananaCount >= 2 {
			set += 1
			pearCount = pearCount - 4
			bananaCount = bananaCount - 2
		} else {
			break
		}
	}

	fmt.Println("set ", set)
	return set
}
