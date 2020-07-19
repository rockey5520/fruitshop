package main

import (
	"fruitshop/controllers"
	"fruitshop/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// initialize DB and load meta data
	models.ConnectDataBase()
	LoadFruitsInventory()

	// Endpoints for customer
	router.GET("/api/v1/customers", controllers.FindCustomers)
	router.GET("/api/v1/customers/:login_id", controllers.FindCustomer)
	router.POST("/api/v1/customers", controllers.CreateCustomer)

	// Endpoints for fruits
	router.GET("/api/v1/fruits", controllers.FindFruits)
	router.GET("/api/v1/fruits/:name", controllers.FindFruit)

	// Endpoints for discounts
	router.GET("/api/v1/discounts/:login_id", controllers.FindDiscounts)

	// Endpoints for cart
	router.GET("/api/v1/cart/:login_id", controllers.FindCart)
	router.POST("/api/v1/cartitem/:login_id", controllers.CreateUpdateItemInCart)
	router.GET("/api/v1/cartitem/:login_id", controllers.GetAllCartItems)

	// Endpoints for coupon
	router.GET("/api/v1/orangecoupon/:login_id", controllers.ApplyOrangeCoupon)

	// Endpoints for coupon
	router.POST("/api/v1/pay/:login_id", controllers.Pay)

	// Use middleware to serve static pages for the website
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/fruitshop-ui", true)))
	router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	err := router.Run()
	if err != nil {
		panic("Unable to invoke router")
	}

}

//LoadFruitsInventory will load fruits inventory metadata
func LoadFruitsInventory() {

	apple := models.Fruit{Name: "Apple", Price: 1.0}
	banana := models.Fruit{Name: "Banana", Price: 1.0}
	pear := models.Fruit{Name: "Pear", Price: 1.0}
	orange := models.Fruit{Name: "Orange", Price: 1.0}

	if err := models.DB.Create(&apple).Error; err != nil {
		panic("Unable to create data into fruits inventory")
	}

	if err := models.DB.Create(&banana).Error; err != nil {
		panic("Unable to create data into fruits inventory")
	}

	if err := models.DB.Create(&pear).Error; err != nil {
		panic("Unable to create data into fruits inventory")
	}

	if err := models.DB.Create(&orange).Error; err != nil {
		panic("Unable to create data into fruits inventory")
	}
}
