package models

/* import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
) */

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	database, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.LogMode(true)

	database.AutoMigrate(&Fruit{},
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

	return database
}

//var DB *gorm.DB

/* func ConnectDataBase() {

	database, err := gorm.Open("sqlite3", "fruitshop.db")

	if err != nil {
		panic("Failed to connect to database!")

	}
	database.LogMode(true)

	database.AutoMigrate(&Fruit{},
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

	DB = database

} */
