package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"fruitshop/api/controllers"
	"fruitshop/api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

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
	if TestDbDriver == "sqlite3" {
		//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		testDbName := os.Getenv("TestDbName")
		server.DB, err = gorm.Open(TestDbDriver, testDbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
		server.DB.LogMode(true)
	}
}

func refreshUserTable() error {

	err := server.DB.DropTableIfExists(&models.Post{}, &models.User{}, &models.Customer{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Customer{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed table(s)")
	return nil
}

func refreshCustomerTable() error {
	err := server.DB.Debug().DropTableIfExists(&models.Customer{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.Customer{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed customer table")
	log.Printf("refreshCustomerTable routine OK !!!")
	return nil
}

func refreshDiscountsTable() error {
	err := server.DB.Debug().DropTableIfExists(
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemCoupon{},
		&models.AppliedSingleItemDiscount{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed discounts table")
	log.Printf("refreshDiscountsTable routine OK !!!")
	return nil
}

func refreshFruitTable() error {
	err := server.DB.Debug().DropTableIfExists(&models.Fruit{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.Fruit{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed Fruit table")
	log.Printf("refreshFruitTable routine OK !!!")
	return nil
}

func refreshCartTable() error {
	err := server.DB.Debug().DropTableIfExists(&models.Cart{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.Cart{},
		&models.Payment{},
		&models.AppliedDualItemDiscount{},
		&models.AppliedSingleItemDiscount{},
		&models.AppliedSingleItemCoupon{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed cart table")
	log.Printf("refreshCartTable routine OK !!!")
	return nil
}

func refreshCartItemTable() error {
	err := server.DB.Debug().DropTableIfExists(&models.CartItem{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.CartItem{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed CartItem table")
	log.Printf("refreshCartItemTable routine OK !!!")
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

	err := server.DB.Debug().Model(&models.Cart{}).Create(&newCart).Error
	if err != nil {
		log.Fatalf("cannot seed Cart table: %v", err)
	}

	log.Printf("seedOneCart routine OK !!!")
	return newCart, nil
}
func seedSingleItemDiscount() (models.AppliedSingleItemDiscount, error) {

	_ = refreshDiscountsTable()

	newDiscount := models.AppliedSingleItemDiscount{
		CartID:  1,
		Savings: 2.0,
	}

	newAppleDiscount := models.SingleItemDiscount{
		FruitID: 1,
		Count:   7,
		Name:    "APPLE10",
		Model: gorm.Model{
			ID: 1,
		},
	}
	err := server.DB.Debug().Model(&models.SingleItemDiscount{}).Create(&newAppleDiscount).Error
	if err != nil {
		log.Fatalf("cannot seed Single item discount table: %v", err)
	}
	err = server.DB.Debug().Model(&models.AppliedSingleItemDiscount{}).Create(&newDiscount).Error
	if err != nil {
		log.Fatalf("cannot seed Single item discount table: %v", err)
	}

	log.Printf("seedSingleItemDiscount routine OK !!!")
	return newDiscount, nil
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

	err := server.DB.Debug().Model(&models.CartItem{}).Create(&newCartItem).Error
	if err != nil {
		log.Fatalf("cannot seed CartItem table: %v", err)
	}

	log.Printf("seedOneCartItem routine OK !!!")
	return newCartItem, nil
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
		LoginID:   "a",
		Cart:      newcart,
	}

	err := server.DB.Debug().Model(&models.Customer{}).Create(&customer).Error
	if err != nil {
		log.Fatalf("cannot seed customers table: %v", err)
	}

	log.Printf("seedOneCustomer routine OK !!!")
	return customer, nil
}

func seedOneUser() (models.User, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Nickname: "Pet",
		Email:    "pet@gmail.com",
		Password: "password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}
	users := []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Kenny Morris",
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func seedFruits() ([]models.Fruit, error) {

	var err error
	if err != nil {
		return nil, err
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
		err := server.DB.Model(&models.User{}).Create(&fruits[i]).Error
		if err != nil {
			return []models.Fruit{}, err
		}
	}
	return fruits, nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.User{
		Nickname: "Sam Phil",
		Email:    "sam@gmail.com",
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}
	var users = []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Magu Frank",
			Email:    "magu@gmail.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}
