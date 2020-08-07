package api

import (
	"fruitshop/api/controllers"
)

var server = controllers.Server{}

func init() {
	// // loads values from .env into the system
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("sad .env file found")
	// }
}

//Run is
func Run() {

	// var err error
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// } else {
	// 	fmt.Println("We are getting the env values")
	// }

	server.InitializeServer("sqlite3", "fruitshop.sqlite")

	server.Run(":8080")

}
