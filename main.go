package main

import "fruitshop/server/api"

func main() {
	// router := gin.Default()
	// initSwagger(router)

	// // initialize DB and load meta data
	// //time.Sleep(10 * time.Second)
	// db := models.SetupModels()

	// // Provide db variable to controllers
	// router.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// 	c.Next()
	// })

	// // Endpoints for customer
	// router.GET("/server/api/v1/customers", controllers.FindCustomers)
	// router.GET("/server/api/v1/customers/:login_id", controllers.FindCustomer)
	// router.POST("/server/api/v1/customers", controllers.CreateCustomer)

	// // Endpoints for fruits
	// router.GET("/server/api/v1/fruits", controllers.FindFruits)
	// router.GET("/server/api/v1/fruits/:name", controllers.FindFruit)

	// // Endpoints for cart
	// router.POST("/server/api/v1/cartitem", controllers.CreateUpdateItemInCart)
	// router.GET("/server/api/v1/cartitem/:cart_id", controllers.GetAllCartItems)
	// router.GET("/server/api/v1/cart/:cart_id", controllers.FindCart)

	// // Endpoints for discounts
	// router.GET("/server/api/v1/discounts/:cart_id", controllers.FindDiscounts)

	// // Endpoints for coupon
	// router.GET("/server/api/v1/orangecoupon/:cart_id/:fruit_id/", controllers.ApplyTimeSensitiveCoupon)

	// // Endpoints for coupon
	// router.POST("/server/api/v1/pay", controllers.Pay)

	// // Use middleware to serve static pages for the website
	// /* router.Use(static.Serve("/", static.LocalFile("/app/webapp/dist/webapp", true)))
	// router.Use(static.Serve("/download", static.LocalFile("./output", true))) */
	// router.Use(static.Serve("/", static.LocalFile("../frontend/dist/fruitshop-ui", true)))
	// router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	// err := router.Run(":8081")

	// if err != nil {
	// 	panic("Unable to invoke router")
	// }
	api.Run()

}

// func initSwagger(engine *gin.Engine) {
// 	docs.SwaggerInfo.Title = "Fruit Store REST API"
// 	docs.SwaggerInfo.Description = "This API is backend service for the fruit store"
// 	docs.SwaggerInfo.Version = "1.0"
// 	docs.SwaggerInfo.Host = "localhost:8081"
// 	docs.SwaggerInfo.BasePath = "/server/api/v1"
// 	docs.SwaggerInfo.Schemes = []string{"http"}

// 	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
// 	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
// }
