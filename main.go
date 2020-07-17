package main

import (
	"fruitshop/controllers"
	"fruitshop/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()
	LoadFruitsInventory()
	router.GET("/customers", controllers.FindCustomers)
	router.GET("/customers/:id", controllers.FindCustomer)
	router.POST("/customers", controllers.CreateCustomer)

	err := router.Run()
	if err != nil {
		panic("Unable to invoke router")
	}

}

func LoadFruitsInventory() {

	apple := models.Fruit{Name: "Apple", Price: 1.0}
	banana := models.Fruit{Name: "Banana", Price: 1.0}
	pear := models.Fruit{Name: "Pear", Price: 1.0}
	orange := models.Fruit{Name: "orange", Price: 1.0}
	if err := models.DB.Create(&apple).Error; err != nil {
		panic("Unable to create inventory")
	}

	if err := models.DB.Create(&banana).Error; err != nil {
		panic("Unable to create inventory")
	}

	if err := models.DB.Create(&pear).Error; err != nil {
		panic("Unable to create inventory")
	}

	if err := models.DB.Create(&orange).Error; err != nil {
		panic("Unable to create inventory")
	}

}
