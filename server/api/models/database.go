package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

var db *gorm.DB
var err error

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	/* user := getEnv("PG_USER", "hugo")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "8080")
	database := getEnv("PG_DB", "tasks") */

	user := getEnv("PG_USER", "postgres")
	password := getEnv("PG_PASSWORD", "secret")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "tasks")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	fmt.Println("fruit id ", dbinfo)

	db, err := gorm.Open("postgres", dbinfo)
	//"host=db port:5432 user=postgres dbname=postgres password=example sslmode=disable"
	//database, err :=  gorm.Open( "postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=example")

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.LogMode(true)

	db.AutoMigrate(&Fruit{},
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
	// Initialise value

	return db
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
