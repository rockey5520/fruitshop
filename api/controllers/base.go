package controllers

import (
	"fmt"
	"fruitshop/api/models"
	"fruitshop/api/seed"
	"log"
	"net/http"

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
func (server *Server) InitializeServer(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

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
	}
	if Dbdriver == "sqlite3" {
		//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
	}

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
	seed.Load(server.DB)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println()
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
