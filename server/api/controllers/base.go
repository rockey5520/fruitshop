package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Unable to connect to  %s database", Dbdriver)
			log.Fatal("Error:", err)
		} else {
			fmt.Printf("Successfully connected to %s database", Dbdriver)
		}
	}

	server.DB.Debug().db.AutoMigrate(&Fruit{},
		&SingleItemDiscount{},
		&DualItemDiscount{},
		&SingleItemCoupon{},
		&AppliedDualItemDiscount{},
		&AppliedSingleItemCoupon{},
		&AppliedSingleItemDiscount{},
		&Payment{},
		&Cart{},
		&CartItem{},
		&Customer{},
	)

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
