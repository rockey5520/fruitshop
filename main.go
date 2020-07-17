package main

import (
	"fruitshop/controllers"
	"fruitshop/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	router.GET("/customers", controllers.FindCustomers)
	router.POST("/customers", controllers.CreateCustomer)

	err := router.Run()
	if err != nil {
		panic("Unable to invoke router")
	}

}
