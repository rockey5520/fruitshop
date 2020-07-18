package controllers

import (
	"fmt"
	"fruitshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddItemInCart will add users choosen fruits to the cart list
func AddItemInCart(c *gin.Context) {
	// Bind the input payload to schema for validations
	var input models.CartItem
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
	//Create/Update/Delete Cart entry based on the count
	cartItem := models.CartItem{CartID: cart.ID, Name: input.Name}
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
	fruit := models.Fruit{Name: input.Name}
	// recalcuate the payment for the cart
	RecalcualtePayments(customer, cart, cartItem, fruit)

	c.JSON(http.StatusOK, gin.H{"data": cartItem})

}

// RecalcualtePayments recalcuates the payment for the cart
func RecalcualtePayments(customer models.Customer, cart models.Cart, cartItem models.CartItem, fruit models.Fruit) {
	// Recalcualte the payments
	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}
	var totalCost float64
	for _, item := range cartItems {
		models.DB.First(&fruit)
		totalCost += float64(item.Count) * fruit.Price
	}
	models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", totalCost)

	//Applying apple 10% discount
	if fruit.Name == "Apple" && cartItem.Count >= 7 {
		var discount models.Discount
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "APPLE10").Update("status", "APPLIED")
		ApplyApple10Discount(customer, cart, cartItem, fruit)
	} else {
		var discount models.Discount
		models.DB.Model(&discount).Where("customer_id = ? AND name = ?", cart.CustomerId, "APPLE10").Update("status", "NOTAPPLIED")
	}
	//Applying pear banana 30% discount
	ApplyBananaPear30Discount(customer, cart, cartItem, fruit)

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
				discount := float64(sets*2) / float64(100) * float64(30)
				models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", cart.Total-discount)
			} else if cartItem.Name == "Banana" {
				discount := float64(sets*2) / float64(100) * float64(30)
				models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", cart.Total-discount)
			}
		}
		discountUpdate = true
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
	var totalCost float64
	var cartItems []models.CartItem
	if err := models.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		fmt.Println("Error ", err)
	}

	var restCost float64
	var appleCost float64
	for _, item := range cartItems {
		if item.Name != "Apple" {
			var fruit models.Fruit
			models.DB.Where("name = ?", item.Name).First(&fruit)
			restCost += float64(item.Count) * fruit.Price
		} else {
			var fruit models.Fruit
			models.DB.Where("name = ?", item.Name).First(&fruit)
			appleCost += float64(item.Count) * fruit.Price
		}
		discount := ((float64(item.Count) * fruit.Price) / 100) * 10
		totalCost = (appleCost - discount) + restCost
	}
	models.DB.Model(&cart).Where("customer_id = ?", customer.ID).Update("total", totalCost)

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
