package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"fruitshop/api/controllers"
	"fruitshop/api/models"

	"github.com/jinzhu/gorm"
)

var server = controllers.Server{}
var fruitInstance = models.Fruit{}

func TestMain(main *testing.M) {

	Database("sqlite3", "fruitshop.sqlite")

	log.Printf("Before calling m.Run() !!!")
	ret := main.Run()
	log.Printf("After calling m.Run() !!!")
	os.Exit(ret)
}

func Database(Dbdriver, DbName string) {
	var err error
	TestDbDriver := Dbdriver
	if TestDbDriver == "sqlite3" {

		testDbName := os.Getenv("TestDbName")
		server.DB, err = gorm.Open(TestDbDriver, testDbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
	}
	server.DB.LogMode(true)
	/*
		If we every wanted to switch to a different database we can use this switch at variable TestDbDriver reading fron env
		And execute appropriate DB in respective environments, Such as all acceptance tests cant run on sqllite in-memory db
		integration tests and production calls can be switched to mysql or prostgres. Here GORM gives us a very good flexibility
		to switch to multiple databases without changing the code.
		Code is placed at the end of this file
	*/

}

func refreshCartTable() error {
	err := server.DB.DropTableIfExists(&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed cart table")
	return nil
}

func refreshCustomerTable() error {
	err := server.DB.DropTableIfExists(&models.Customer{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.Customer{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed customer table")
	return nil
}

func refreshCartItemTable() error {
	err := server.DB.DropTableIfExists(&models.CartItem{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.CartItem{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed CartItem table")
	return nil
}

func seedOneCart() (models.Cart, error) {

	_ = refreshCartTable()

	newCart := models.Cart{
		CustomerId:   1,
		Total:        5,
		TotalSavings: 2,
		Status:       "OPEN",
	}

	err := server.DB.Model(&models.Cart{}).Create(&newCart).Error
	if err != nil {
		log.Fatalf("cannot seed Cart table: %v", err)
	}

	log.Printf("seedOneCart routine OK !!!")
	return newCart, nil
}

func seedOneCustomer() (models.Customer, error) {

	_ = refreshCustomerTable()
	_ = refreshCartTable()

	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}
	customer := models.Customer{
		FirstName: "Rakesh",
		LastName:  "Mothukuri",
		LoginID:   "rockey5520",
		Cart:      newcart,
	}

	err := server.DB.Model(&models.Customer{}).Create(&customer).Error
	if err != nil {
		log.Fatalf("cannot seed customers table: %v", err)
	}

	log.Printf("seedOneCustomer routine OK !!!")
	return customer, nil
}

func seedOneCartItem() (models.CartItem, error) {

	_ = refreshCartItemTable()

	newCartItem := models.CartItem{
		CartID:              1,
		FruitID:             1,
		Name:                "Apple",
		Quantity:            10,
		ItemTotal:           10,
		ItemDiscountedTotal: 0.0,
	}

	err := server.DB.Model(&models.CartItem{}).Create(&newCartItem).Error
	if err != nil {
		log.Fatalf("cannot seed CartItem table: %v", err)
	}

	log.Printf("seedOneCartItem routine OK !!!")
	return newCartItem, nil
}

func seedFruits() error {

	var err error
	if err != nil {
		return err
	}
	fruits := []models.Fruit{
		models.Fruit{
			Name:  "Apple",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Pear",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Banana",
			Price: 1.0,
		},
		models.Fruit{
			Name:  "Orange",
			Price: 1.0,
		},
	}

	for i, _ := range fruits {
		err := server.DB.Model(&models.Fruit{}).Create(&fruits[i]).Error
		if err != nil {
			return err
		}
	}

	log.Printf("seedUsers routine OK !!!")
	return nil
}

func refreshFruitTable() error {
	err := server.DB.DropTableIfExists(&models.Fruit{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.Fruit{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed Fruit table")
	log.Printf("refreshFruitTable routine OK !!!")
	return nil
}

/*
		// If we every wanted to switch to a different database we can use this switch at variable TestDbDriver reading fron env
		// And execute appropriate DB in respective environments, Such as all acceptance tests cant run on sqllite in-memory db
		// integration tests and production calls can be switched to mysql or prostgres.
	    // Here GORM gives us a very good flexibility to switch to multiple databases without changing the code.
				TestDbDriver := os.Getenv("TestDbDriver")

			if TestDbDriver == "mysql" {
				DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
				server.DB, err = gorm.Open(TestDbDriver, DBURL)
				if err != nil {
					fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
					log.Fatal("This is the error:", err)
				} else {
					fmt.Printf("We are connected to the %s database\n", TestDbDriver)
				}
			}
			if TestDbDriver == "postgres" {
				DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
				server.DB, err = gorm.Open(TestDbDriver, DBURL)
				if err != nil {
					fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
					log.Fatal("This is the error:", err)
				} else {
					fmt.Printf("We are connected to the %s database\n", TestDbDriver)
				}
			}
*/
