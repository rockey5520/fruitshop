package controllers

import (
	"fmt"
	"fruitshop/api/models"
	"fruitshop/api/seed"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// InitializeServer initializes selected DB from env file( defaults to sqllite3 ) and starts application on port 8080 using gorilla mux
func (server *Server) InitializeServer(Dbdriver, DbName string) {
	var err error
	/*
		        If we every wanted to switch to a different database we can use this switch at variable TestDbDriver reading fron env
				And execute appropriate DB in respective environments, Such as all acceptance tests cant run on sqllite in-memory db
				integration tests and production calls can be switched to mysql or prostgres. Here GORM gives us a very good flexibility
				to switch to multiple databases without changing the code. I have wriiten the code for it and placed below for reference.
	*/
	if Dbdriver == "sqlite3" {
		server.DB, err = gorm.Open(Dbdriver, DbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
	}
	// only for debugging purpose and we need to turn it off in production
	server.DB.LogMode(true)

	//database migration
	server.DB.DropTableIfExists(
		&models.Customer{},
		&models.Cart{},
		&models.CartItem{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{},
		&models.Payment{},
	)

	server.DB.Debug().AutoMigrate(
		&models.Customer{},
		&models.Cart{},
		&models.CartItem{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{},
		&models.Payment{},
	)

	server.Router = mux.NewRouter()
	seed.LoadTablesAndBusinessRules(server.DB)
	server.initializeRoutes()

}

func (server *Server) Run(addr string, start time.Time) {
	fmt.Println()
	elapsed := time.Since(start)
	log.Printf("Application took %s to start", elapsed)
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))

}

/*
 if Dbdriver == "mysql" {
 	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
 	server.DB, err = gorm.Open(Dbdriver, DBURL)
 	if err != nil {
 		fmt.Printf("Cannot connect to %s database", Dbdriver)
 		log.Fatal("This is the error:", err)
 	} else {
 		fmt.Printf("We are connected to the %s database", Dbdriver)
 	}
 }
 if Dbdriver == "postgres" {
 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
 	server.DB, err = gorm.Open(Dbdriver, DBURL)
 	if err != nil {
 		fmt.Printf("Cannot connect to %s database", Dbdriver)
 		log.Fatal("This is the error:", err)
 	} else {
 		fmt.Printf("We are connected to the %s database", Dbdriver)
 	}
 }*/
