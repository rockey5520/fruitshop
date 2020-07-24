package main

import (
	"fruitshop/controllers"
	"fruitshop/docs"
	"fruitshop/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	initSwagger(router)

	// initialize DB and load meta data
	//time.Sleep(10 * time.Second)
	db := models.SetupModels()
	loadFruitsAndDiscountsTableMetaData(db)

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Endpoints for customer
	router.GET("/api/v1/customers", controllers.FindCustomers)
	router.GET("/api/v1/customers/:login_id", controllers.FindCustomer)
	router.POST("/api/v1/customers", controllers.CreateCustomer)

	// Endpoints for fruits
	router.GET("/api/v1/fruits", controllers.FindFruits)
	router.GET("/api/v1/fruits/:name", controllers.FindFruit)

	// Endpoints for cart
	router.POST("/api/v1/cartitem", controllers.CreateUpdateItemInCart)
	router.GET("/api/v1/cartitem/:cart_id", controllers.GetAllCartItems)
	router.GET("/api/v1/cart/:cart_id", controllers.FindCart)

	// Endpoints for discounts
	router.GET("/api/v1/discounts/:cart_id", controllers.FindDiscounts)

	// Endpoints for coupon
	router.GET("/api/v1/orangecoupon/:cart_id/:fruit_id/", controllers.ApplyTimeSensitiveCoupon)

	// Endpoints for coupon
	router.POST("/api/v1/pay", controllers.Pay)

	// Use middleware to serve static pages for the website
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/fruitshop-ui", true)))
	router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	err := router.Run(":8082")

	if err != nil {
		panic("Unable to invoke router")
	}

}

func initSwagger(engine *gin.Engine) {
	docs.SwaggerInfo.Title = "Fruit Store REST API"
	docs.SwaggerInfo.Description = "This API is backend service for the fruit store"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8082"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	url := ginSwagger.URL("http://localhost:8082/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func loadFruitsAndDiscountsTableMetaData(db *gorm.DB) {

	appleItemDiscount := models.SingleItemDiscount{Count: 7, Discount: 10, Name: "APPLE10"}
	orangeSingleItemCoupon := models.SingleItemCoupon{
		Discount: 30,
		Name:     "ORANGE30",
	}
	apple := models.Fruit{
		Name: "Apple",
		SingleItemDiscount: []models.SingleItemDiscount{
			appleItemDiscount,
		},
		Price: 1,
	}
	banana := models.Fruit{
		Name:  "Banana",
		Price: 1,
	}
	pear := models.Fruit{
		Name:  "Pear",
		Price: 1,
	}
	orange := models.Fruit{
		Name: "Orange",
		SingleItemCoupon: []models.SingleItemCoupon{
			orangeSingleItemCoupon,
		},
		Price: 1,
	}

	if err := db.Create(&apple).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&banana).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&pear).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&orange).Error; err != nil {
		panic("Unable to create fruit inventory")
	}

	var pearFromDB models.Fruit
	db.Where("name = ?", "Pear").First(&pearFromDB)
	var bananaFromDB models.Fruit
	db.Where("name = ?", "Banana").First(&bananaFromDB)
	dualItemDiscount := models.DualItemDiscount{
		FruitID:   pearFromDB.ID,
		FruitID_1: pearFromDB.ID,
		FruitID_2: bananaFromDB.ID,
		Count_1:   4,
		Count_2:   2,
		Name:      "PEARBANANA30",
		Discount:  30,
	}
	if err := db.Create(&dualItemDiscount).Error; err != nil {
		panic("Unable to create Single item discount inventory")
	}

}
