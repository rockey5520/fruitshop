package main

import (
	"fruitshop/controllers"
	"fruitshop/docs"
	"fruitshop/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	initSwagger(router)

	// initialize DB and load meta data
	models.ConnectDataBase()
	LoadFruitsInventory()

	// Endpoints for customer
	router.GET("/server/api/v1/customers", controllers.FindCustomers)
	router.GET("/server/api/v1/customers/:login_id", controllers.FindCustomer)
	router.POST("/server/api/v1/customers", controllers.CreateCustomer)

	// Endpoints for fruits
	router.GET("/server/api/v1/fruits", controllers.FindFruits)
	router.GET("/server/api/v1/fruits/:name", controllers.FindFruit)

	// Endpoints for discounts
	router.GET("/server/api/v1/discounts/:login_id", controllers.FindDiscounts)

	// Endpoints for cart
	router.GET("/server/api/v1/cart/:login_id", controllers.FindCart)
	router.POST("/server/api/v1/cartitem/:login_id", controllers.CreateUpdateItemInCart)
	router.GET("/server/api/v1/cartitem/:login_id", controllers.GetAllCartItems)

	// Endpoints for coupon
	router.GET("/server/api/v1/orangecoupon/:login_id", controllers.ApplyOrangeCoupon)

	// Endpoints for coupon
	router.POST("/server/api/v1/pay/:login_id", controllers.Pay)

	// Use middleware to serve static pages for the website
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/fruitshop-ui", true)))
	router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	err := router.Run()
	if err != nil {
		panic("Unable to invoke router")
	}

}

func initSwagger(engine *gin.Engine) {
	docs.SwaggerInfo.Title = "Fruit Store REST API"
	docs.SwaggerInfo.Description = "This API is backend service for the fruit store"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
