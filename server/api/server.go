package api

import (
	"fruitshop/api/controllers"
)

var server = controllers.Server{}

//Run function will initialize server on port 8080 and spin up the DB to store the application data.
func Run() {
	/*
		We can always use the below commented code to load environment variable from env file and send appropriate DB params
		to the InitializeServer function to load right DB for each use case( such as unit testing, acceptance testing, Integration testing and production)
	*/
	server.InitializeServer("sqlite3", "fruitshop.sqlite")
	server.Run(":8080")
}

/*
	// possibility to load environment varables from a .env file and pass then to Database() function
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}*/
