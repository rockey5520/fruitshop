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
	router.GET("/server/api/v1/customers", controllers.FindCustomers)
	router.GET("/server/api/v1/customers/:login_id", controllers.FindCustomer)
	router.POST("/server/api/v1/customers", controllers.CreateCustomer)

	// Endpoints for fruits
	router.GET("/server/api/v1/fruits", controllers.FindFruits)
	router.GET("/server/api/v1/fruits/:id", controllers.FindFruit)

	// Endpoints for discounts
	router.GET("/server/api/v1/discounts/:login_id", controllers.FindDiscounts)

	// Use middleware to serve static
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/fruitshop-ui", true)))
	router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	err := router.Run()
	if err != nil {
		panic("Unable to invoke router")
	}

}

func cleanDB() {
	var fruits models.Fruit
	models.DB.Where("id > 0").Find(&fruits).Unscoped().Delete(&fruits)
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
