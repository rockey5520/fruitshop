package fruitshop

import (
	"fmt"
	cart "fruitshop/gen/cart"
	fruit "fruitshop/gen/fruit"
	payment "fruitshop/gen/payment"
	"fruitshop/gen/user"
	"math"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User *user.UserManagement
type Fruit *fruit.FruitManagement
type Cart *cart.CartManagement
type Payment *payment.PaymentManagement

// InitDB is the function that starts a database file and table structures
// if not created then returns db object for next functions
func InitDB() *gorm.DB {
	// Opening file

	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the user table if it doesn't exist
	var TableStruct = user.UserManagement{}
	if !db.HasTable(TableStruct) {
		db.CreateTable(TableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(TableStruct)
	}
	// Creating fruit table if it doesn't exist
	var FruitTableStruct = fruit.FruitManagement{}
	if !db.HasTable(FruitTableStruct) {
		db.CreateTable(FruitTableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(FruitTableStruct)
	}

	var CartTableStruct = cart.CartManagement{}
	if !db.HasTable(CartTableStruct) {
		db.CreateTable(CartTableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(CartTableStruct)
	}

	var PaymentTableStruct = payment.PaymentManagement{}
	if !db.HasTable(PaymentTableStruct) {
		db.CreateTable(PaymentTableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(PaymentTableStruct)
	}

	return db
}

// GetClient retrieves one client by its ID
func GetUser(ID string) (user.UserManagement, error) {
	db := InitDB()
	defer db.Close()

	var users user.UserManagement
	db.Where("ID = ?", ID).First(&users)

	return users, err
}

// CreateClient created a client row in DB
func CreateUser(user User) error {
	db := InitDB()
	defer db.Close()
	err := db.Create(&user).Error
	return err
}

// ListClients retrieves the clients stored in Database
func ListUsers() (user.UserManagementCollection, error) {
	db := InitDB()
	defer db.Close()
	var users user.UserManagementCollection
	err := db.Find(&users).Error
	return users, err
}

// CreateClient created a fruit row in DB
func CreateFruit(fruit Fruit) error {
	db := InitDB()
	defer db.Close()
	err := db.Create(&fruit).Error
	return err
}

// ListClients retrieves the fruits stored in Database
var counter = 0

func ListFruits() (fruit.FruitManagementCollection, error) {
	db := InitDB()
	defer db.Close()
	var fruits fruit.FruitManagementCollection
	var count int
	db.Where("name = ?", "Apple").Find(&fruits).Count(&count)
	if count == 0 {
		initialize()
	}
	counter++

	err := db.Find(&fruits).Error
	return fruits, err
}

// CreateClient created a client row in DB
func CreateCart(user User) error {
	db := InitDB()
	defer db.Close()
	paymentID := user.ID + user.ID
	paymentStatus := "NOTPAID"
	paymentAmount := float64(0)
	cartItemApple := cart.CartManagement{
		CartID:      user.ID,
		Name:        "Apple",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemBanana := cart.CartManagement{
		CartID:      user.ID,
		Name:        "Banana",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemPear := cart.CartManagement{
		CartID:      user.ID,
		Name:        "Pear",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemOrange := cart.CartManagement{
		CartID:      user.ID,
		Name:        "Orange",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	payment := payment.PaymentManagement{
		ID:            &paymentID,
		CartID:        user.ID,
		PaymentStatus: &paymentStatus,
		Amount:        &paymentAmount,
	}

	err := db.Create(&cartItemApple).Error
	db.Create(&cartItemBanana)
	db.Create(&cartItemPear)
	db.Create(&cartItemOrange)
	db.Create(&payment)

	return err
}

// AddItemInCart updated a cart entry row in DB
func AddItemInCart(cart Cart) error {
	db := InitDB()
	defer db.Close()
	var fruits fruit.FruitManagement
	db.Where("Name = ?", cart.Name).Find(&fruits)
	//db.Table("fruit_managements").Where("Name = ?", cart.Name).First(&fruits)
	err := db.Model(&cart).
		Where("name = ?", cart.Name).
		Update("count", cart.Count).
		Update("cost_per_item", fruits.Cost).
		Update("total_cost", fruits.Cost*float64(cart.Count)).
		Error

	paymentId := cart.CartID + cart.CartID
	amount := fruits.Cost * float64(cart.Count)
	var payments payment.PaymentManagement
	db.Where("ID = ?", paymentId).First(&payments)
	payment := &payment.PaymentManagement{
		ID:     &paymentId,
		CartID: cart.CartID,
	}
	var currentTotal float64 = *payments.Amount + amount
	db.Find(&payment).
		Update("amount", currentTotal).
		Update("PaymentStatus", "NOTPAID")

	return err
}

func getCurrentCost(cart_id string, name string) float64 {
	db := InitDB()
	defer db.Close()

	var carts cart.CartManagement
	db.Where("Name = ?", name).First(&carts)
	fmt.Println(carts)
	return carts.TotalCost
}

// RemoveItemInCart creates a cart entry row in DB
func RemoveItemInCart(cart Cart) error {
	db := InitDB()
	defer db.Close()

	var fruits fruit.FruitManagement
	db.Table("fruit_managements").Where("name = ?", cart.Name).First(&fruits)

	paymentId := cart.CartID + cart.CartID
	amount := fruits.Cost * getCurrentCost(cart.CartID, cart.Name)

	var payments payment.PaymentManagement
	db.Where("ID = ?", paymentId).First(&payments)
	payment := &payment.PaymentManagement{
		ID:     &paymentId,
		CartID: cart.CartID,
	}

	var currentTotal float64 = math.Abs(*payments.Amount - amount)

	if currentTotal == 0 {
		db.Model((&payment)).
			Update("PaymentStatus", "PAID")
	}

	db.Model((&payment)).
		Update("amount", currentTotal)
	err := db.Model(&cart).
		Where("name = ?", cart.Name).
		Update("count", cart.Count).
		Update("cost_per_item", fruits.Cost).
		Update("total_cost", fruits.Cost*float64(cart.Count)).
		Error

	return err
}

func ListAllItemsInCartForId(CartID string) (cart.CartManagementCollection, error) {
	db := InitDB()
	defer db.Close()
	var carts cart.CartManagementCollection
	err := db.Find(&carts).Where("CartID = ?", CartID).Error
	return carts, err
}

// CreateCartItem creates a cart entry row in DB
func GetPaymentAmoutFromCart(input Payment) (payment.PaymentManagement, error) {
	db := InitDB()
	defer db.Close()
	var totalAmount float64
	var cart cart.CartManagementCollection
	db.Table("cart_managements").Where("cart_id = ?", input.CartID).Find(&cart)

	for _, x := range cart {
		totalAmount += float64(x.TotalCost)
	}
	var payments payment.PaymentManagement
	db.Model(&payments).Where("ID = ?", input.ID).Update("amount", totalAmount)

	db.Where("ID = ?", input.ID).First(&payments)

	fmt.Println("totalAmount", totalAmount)

	return payments, err
}

func PayAmount(input Payment) (payment.PaymentManagement, error) {
	db := InitDB()
	defer db.Close()

	var payments payment.PaymentManagement
	db.Model(&payments).
		Where("ID = ?", input.ID).
		Update("amount", 0).
		Update("PaymentStatus", "PAID")

	user := user.UserManagement{
		ID: input.CartID,
	}
	paymentID := user.ID + user.ID
	paymentStatus := "NOTPAID"
	paymentAmount := float64(0)
	cartItemApple := cart.CartManagement{
		CartID: user.ID,
		Name:   "Apple",
	}
	cartItemBanana := cart.CartManagement{
		CartID: user.ID,
		Name:   "Banana",
	}
	cartItemPear := cart.CartManagement{
		CartID: user.ID,
		Name:   "Pear",
	}
	cartItemOrange := cart.CartManagement{
		CartID: user.ID,
		Name:   "Orange",
	}
	payment := payment.PaymentManagement{
		ID:            &paymentID,
		CartID:        user.ID,
		PaymentStatus: &paymentStatus,
		Amount:        &paymentAmount,
	}

	err := db.Model(&cartItemApple).
		Update("Count", 0).Error
	db.Model(&cartItemBanana).
		Update("Count", 0)
	db.Model(&cartItemPear).
		Update("Count", 0)
	db.Model(&cartItemOrange).
		Update("Count", 0)
	db.Model(&payment).
		Update("PaymentStatus", "PAID").
		Update("Amount", 0)
	return payments, err
}

func initialize() {
	db := InitDB()
	defer db.Close()
	apple := fruit.FruitManagement{
		Name: "Apple",
		Cost: 1.0,
	}
	banana := fruit.FruitManagement{
		Name: "Banana",
		Cost: 1,
	}
	pear := fruit.FruitManagement{
		Name: "Pear",
		Cost: 1,
	}
	orange := fruit.FruitManagement{
		Name: "Orange",
		Cost: 1,
	}
	db.NewRecord(apple)
	db.Create(&apple)
	db.NewRecord(banana)
	db.Create(&banana)
	db.NewRecord(pear)
	db.Create(&pear)
	db.NewRecord(orange)
	db.Create(&orange)
}
