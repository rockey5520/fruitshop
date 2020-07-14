package fruitshop

import (
	"fmt"
	cart "fruitshop/gen/cart"
	"fruitshop/gen/coupon"
	"fruitshop/gen/discount"
	fruit "fruitshop/gen/fruit"
	payment "fruitshop/gen/payment"
	"fruitshop/gen/user"
	"math"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User *user.UserManagement
type Fruit *fruit.FruitManagement
type Cart *cart.CartManagement
type Payment *payment.PaymentManagement
type Discount *discount.DiscountManagement
type Coupon *coupon.CouponManagement

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

	var DiscountTableStruct = discount.DiscountManagement{}
	if !db.HasTable(DiscountTableStruct) {
		db.CreateTable(DiscountTableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(DiscountTableStruct)
	}

	var CouponTableStruct = coupon.CouponManagement{}
	if !db.HasTable(CouponTableStruct) {
		db.CreateTable(CouponTableStruct)
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(CouponTableStruct)
	}

	return db
}

// GetClient retrieves one client by its ID
func GetUser(input User) (user.UserManagement, error) {
	db := InitDB()
	defer db.Close()

	var users user.UserManagement
	db.Where("id = ?", input.ID).
		First(&users)

	return users, err
}

// CreateClient created a client row in DB
func CreateUser(user User) (user.UserManagement, error) {
	db := InitDB()
	defer db.Close()
	var err1 error
	var err2 error
	var err3 error
	var err4 error
	// If user does not exists then create a user
	if !doesUserExist(user) {
		err1 = db.Create(&user).Error
		err2 = CreateCart(user)
		err3 = CreateDiscount(user)
		err4 = CreateCoupon(user)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("An error occurred...")
			fmt.Println(err)

		}
	}
	result, err := GetUser(user)
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}
	return result, err
}

func setCouponOrange30Active(input Coupon) {

	db := InitDB()
	defer db.Close()

	var coupon coupon.CouponManagement
	db.Model(coupon).
		Where("id = ?", input.ID).
		Update("status", "ACTIVE")

}

func updateCouponStatus(input Coupon) {
	db := InitDB()
	defer db.Close()

	var coupon coupon.CouponManagement
	db.Where("ID = ?", input.ID).First(&coupon)

	user := user.UserManagement{
		UserID: input.UserID,
	}

	var carts cart.CartManagementCollection
	fmt.Println("input.UserID", input.UserID)
	carts, err = ListAllItemsInCartForId(&user)
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}
	for _, x := range carts {
		if x.Name == "Orange" && x.Count > 0 {
			discountCalculated := ((float64(x.Count) * x.CostPerItem) / 100) * 30
			updatedTotalCost := x.TotalCost - discountCalculated
			cart := cart.CartManagement{
				UserID:    input.UserID,
				Name:      x.Name,
				TotalCost: updatedTotalCost,
			}
			err = UpdateItemInCart(&cart)
			if err != nil {
				fmt.Println("An error occurred...")
				fmt.Println(err)

			}
			discount := discount.DiscountManagement{
				UserID: input.UserID,
			}
			db.Model(&discount).Where("user_id = ?", input.UserID).
				Where("name = ?", "ORANGE30").
				Update("status", "APPLIED")

		}
	}

	time.Sleep(10 * time.Second)

	db.Model(coupon).
		Where("id = ?", input.ID).
		Update("status", "NOTACTIVE")

	discount := discount.DiscountManagement{
		UserID: input.UserID,
	}
	db.Model(&discount).Where("user_id = ?", input.UserID).
		Where("name = ?", "ORANGE30").
		Update("status", "NOTAPPLIED")

	for _, x := range carts {
		if x.Name == "Orange" && x.Count > 0 {
			updatedTotalCost := float64(x.Count) * x.CostPerItem
			cart := cart.CartManagement{
				UserID:    input.UserID,
				Name:      x.Name,
				TotalCost: updatedTotalCost,
			}
			err = UpdateItemInCart(&cart)
			if err != nil {
				fmt.Println("An error occurred...")
				fmt.Println(err)

			}

			db.Model(&discount).Where("user_id = ?", input.UserID).
				Where("name = ?", "ORANGE30").
				Update("status", "APPLIED")

		}
	}

}

//CreateCoupon creates coupon table and dummy data
func CreateCoupon(input User) error {
	db := InitDB()
	defer db.Close()

	var orange30 string = "ORANGE30"
	var notactive string = "NOTACTIVE"

	coupon := coupon.CouponManagement{
		ID:     &input.ID,
		UserID: input.UserID,
		Status: &notactive,
		Name:   &orange30,
	}

	err := db.Create(&coupon).Error
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}
	return err
}

// doesUserExist checks if user exists
func doesUserExist(input User) bool {
	var res user.UserManagement
	res, err = GetUser(input)
	return res.ID == input.ID
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

func getDiscounts(input Discount) (discount.DiscountManagementCollection, error) {
	db := InitDB()
	defer db.Close()
	var discounts discount.DiscountManagementCollection
	err := db.Where("user_id = ?", input.UserID).Find(&discounts).Error
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}
	return discounts, err
}

// CreateClient created a client row in DB
func CreateCart(user User) error {
	db := InitDB()
	defer db.Close()
	paymentID := user.ID
	paymentStatus := "NOTPAID"
	paymentAmount := float64(0)
	cartItemApple := cart.CartManagement{
		UserID:      user.UserID,
		Name:        "Apple",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemBanana := cart.CartManagement{
		UserID:      user.UserID,
		Name:        "Banana",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemPear := cart.CartManagement{
		UserID:      user.UserID,
		Name:        "Pear",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	cartItemOrange := cart.CartManagement{
		UserID:      user.UserID,
		Name:        "Orange",
		Count:       0,
		CostPerItem: 1,
		TotalCost:   0,
	}
	payment := payment.PaymentManagement{
		ID:            &paymentID,
		UserID:        user.UserID,
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

// CreateDiscount created a discount row in DB
func CreateDiscount(user User) error {
	db := InitDB()
	defer db.Close()

	var apple10 string = "APPLE10"
	var pearbanana string = "PEARBANANA30"
	var orange30 string = "ORANGE30"

	var notapplied string = "NOTAPPLIED"

	discountApple10 := discount.DiscountManagement{
		UserID: user.UserID,
		Name:   &apple10,
		Status: &notapplied,
	}

	discountpearbanana := discount.DiscountManagement{
		UserID: user.UserID,
		Name:   &pearbanana,
		Status: &notapplied,
	}
	discountorange30 := discount.DiscountManagement{
		UserID: user.UserID,
		Name:   &orange30,
		Status: &notapplied,
	}

	err := db.Create(&discountApple10).Error
	db.Create(&discountpearbanana)
	db.Create(&discountorange30)
	return err
}

// AddItemInCart updated a cart entry row in DB
func AddItemInCart(cart Cart) error {
	db := InitDB()
	defer db.Close()
	var fruits fruit.FruitManagement
	db.Where("Name = ?", cart.Name).Find(&fruits)

	err := db.Model(&cart).
		Where("user_id = ?", cart.UserID).
		Where("name = ?", cart.Name).
		Update("count", cart.Count).
		Update("cost_per_item", fruits.Cost).
		Update("total_cost", fruits.Cost*float64(cart.Count)).
		Error
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}
	// check the dicounts can be applied

	if cart.Name == "Apple" {
		checkApple10(cart)
	}
	if cart.Name == "Banana" || cart.Name == "Pear" {
		checkPearBanana30(cart)
	}

	var payments payment.PaymentManagement
	db.Where("ID = ?", cart.UserID+cart.UserID).
		Where("user_id = ?", cart.UserID).
		First(&payments)

	user := user.UserManagement{
		ID:     cart.UserID + cart.UserID,
		UserID: cart.UserID,
	}

	err = updatePayment(&user)
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)

	}

	return err
}

func updatePayment(input User) error {
	db := InitDB()
	defer db.Close()
	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(input)
	var totalAmount float64
	for _, x := range carts {
		totalAmount += x.TotalCost
	}

	var payment payment.PaymentManagement
	fmt.Println("total amount ", totalAmount)
	err := db.Where("user_id = ?", input.UserID).
		Where("id = ?", input.ID).Find(&payment).
		Update("amount", totalAmount).Error

	/* err := db.Find(payment).
	Update("amount", totalAmount).
	Update("payment_status", "NOTPAID").Error */
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}

	return err

}

func ListAllItemsInCartForId(input User) (cart.CartManagementCollection, error) {
	db := InitDB()
	defer db.Close()
	var carts cart.CartManagementCollection
	err := db.Where("user_id = ?", input.UserID).Find(&carts).Error
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}
	return carts, err
}

func checkApple10(input Cart) {
	fmt.Println("checkapple10 - 2")
	user := user.UserManagement{
		UserID: input.UserID,
	}
	var apple10 string = "APPLE10"
	var applied string = "APPLIED"
	var notapplied string = "NOTAPPLIED"

	var discountUpdate bool
	discount := discount.DiscountManagement{
		UserID: input.UserID,
		Name:   &apple10,
	}

	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(&user)
	fmt.Println("checkapple10 - 3")
	for _, x := range carts {
		cart := x
		fmt.Println(cart.Name)
		fmt.Println(cart.Count)
		fmt.Println(cart.Name == "Apple")
		if cart.Name == "Apple" {
			fmt.Println("checkapple10 - 4")
			if cart.Count >= 7 {
				fmt.Println("cart.Count ", cart.Count)
				discount := ((float64(cart.Count) * cart.CostPerItem) / 100) * 10
				cart.TotalCost = cart.TotalCost - discount
				err = UpdateItemInCart(cart)
				discountUpdate = true
				if err != nil {
					fmt.Println("An error occurred...")
					fmt.Println(err)
				}

			}
		}
	}
	if discountUpdate {
		discount.Status = &applied
		err = UpdateDiscount(&discount)
		if err != nil {
			fmt.Println("An error occurred...")
			fmt.Println(err)
		}
	} else {
		discount.Status = &notapplied
		err = UpdateDiscount(&discount)
		if err != nil {
			fmt.Println("An error occurred...")
			fmt.Println(err)
		}
	}

}

func checkPearBanana30(input Cart) {

	user := user.UserManagement{
		UserID: input.UserID,
	}
	var pearbanana30 string = "PEARBANANA30"
	var applied string = "APPLIED"
	var notapplied string = "NOTAPPLIED"
	var pearCount int
	var bananaCount int

	var discountUpdate bool
	discount := discount.DiscountManagement{
		UserID: input.UserID,
		Name:   &pearbanana30,
	}

	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(&user)
	for _, x := range carts {
		cart := x
		if cart.Name == "Pear" {
			pearCount = cart.Count

		} else if cart.Name == "Banana" {
			bananaCount = cart.Count
		}
	}

	fmt.Println("pearCount ", pearCount)
	fmt.Println("bananaCount ", bananaCount)
	sets := getSets(pearCount, bananaCount)
	if sets != 0 {
		for _, x := range carts {
			cart := x

			if cart.Name == "Pear" {
				discount := float64(sets*2) / float64(100) * float64(30)
				cart.TotalCost = float64(cart.Count)*cart.CostPerItem - discount
				fmt.Println("cart.Name", cart.Name)
				fmt.Println("cart.TotalCost", cart.TotalCost)
				fmt.Println("discount", discount)
				updateCartTotalCostByID(cart)
			} else if cart.Name == "Banana" {

				discount := float64(sets*2) / float64(100) * float64(30)
				cart.TotalCost = float64(cart.Count)*cart.CostPerItem - discount
				fmt.Println("cart.Name", cart.Name)
				fmt.Println("cart.TotalCost", cart.TotalCost)
				fmt.Println("discount", discount)
				updateCartTotalCostByID(cart)
			}

		}
		discountUpdate = true
	} else {
		recalculateCartTable(input)
		discountUpdate = false
	}

	if discountUpdate {
		discount.Status = &applied
		err = UpdateDiscount(&discount)
		if err != nil {
			fmt.Println("An error occurred...")
			fmt.Println(err)
		}
	} else {
		discount.Status = &notapplied
		err = UpdateDiscount(&discount)
		if err != nil {
			fmt.Println("An error occurred...")
			fmt.Println(err)
		}
	}

}

func updateCartTotalCostByID(input Cart) {
	db := InitDB()
	defer db.Close()
	//checkApple10(input)
	err := db.Model(&input).
		Where("user_id = ?", input.UserID).
		Where("name = ?", input.Name).
		Update("total_cost", input.TotalCost).
		Error

	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}

}

func updateDiscountCartAndPayment(input Cart) {
	db := InitDB()
	defer db.Close()

	err := db.Model(&input).
		Where("user_id = ?", input.UserID).
		Where("name = ?", input.Name).
		Update("total_cost", input.TotalCost).
		Error

	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}

	updatePayments(input)

}

func recalculateCartTable(input Cart) {
	db := InitDB()
	defer db.Close()

	user := user.UserManagement{
		UserID: input.UserID,
	}

	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(&user)

	for _, x := range carts {
		totalCost := x.CostPerItem * float64(x.Count)
		x.TotalCost = totalCost
		updateCartTotalCostByID(x)
	}

}

func updatePayments(input Cart) {
	user := user.UserManagement{
		UserID: input.UserID,
	}
	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(&user)
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}
	var totalCartCost float64

	for _, i := range carts {
		totalCartCost += i.TotalCost
	}

	ID := input.UserID + input.UserID
	payment := payment.PaymentManagement{
		ID:     &ID,
		UserID: input.UserID,
	}

	if totalCartCost <= 0 {
		db.Find(&payment).
			Update("amount", 0).
			Update("PaymentStatus", "PAID")
	} else {
		db.Find(&payment).
			Update("amount", totalCartCost).
			Update("PaymentStatus", "NOTPAID")
	}
}

func getSets(pear int, banana int) int {

	pearCount := pear
	bananaCount := banana
	var set int

	for i := 0; i < pearCount; i++ {
		if pearCount >= 4 && bananaCount >= 2 {
			set += 1
			pearCount = pearCount - 4
			bananaCount = bananaCount - 2
		} else {
			break
		}
	}

	fmt.Println("set ", set)
	return set
}

// UpdateDiscount updated a cart entry row in DB
func UpdateDiscount(input Discount) error {
	db := InitDB()
	defer db.Close()
	err := db.Model(&input).
		Where("user_id = ?", input.UserID).
		Where("name = ?", input.Name).
		Update("status", input.Status).Error

	return err
}

// UpdateItemInCart updated a cart entry row in DB
func UpdateItemInCart(input Cart) error {
	db := InitDB()
	defer db.Close()
	var fruits fruit.FruitManagement
	db.Where("Name = ?", input.Name).Find(&fruits)
	ID := input.UserID + input.UserID
	err := db.Model(&input).
		Where("user_id = ?", input.UserID).
		Where("name = ?", input.Name).
		Update("total_cost", input.TotalCost).
		Error
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}
	user := user.UserManagement{
		UserID: input.UserID,
	}

	var carts cart.CartManagementCollection
	carts, err = ListAllItemsInCartForId(&user)
	var totalCartCost float64

	for _, i := range carts {
		totalCartCost += i.TotalCost
	}

	payment := payment.PaymentManagement{
		ID:     &ID,
		UserID: input.UserID,
	}

	if totalCartCost <= 0 {
		db.Find(&payment).
			Update("amount", 0).
			Update("PaymentStatus", "PAID")
	} else {
		db.Find(&payment).
			Update("amount", totalCartCost).
			Update("PaymentStatus", "NOTPAID")
	}

	return err
}

func getCurrentCost(ID string, user_id string, name string) float64 {
	db := InitDB()
	defer db.Close()

	var carts cart.CartManagement
	db.Where("Name = ?", name).
		Where("ID = ?", ID).
		First(&carts)
	fmt.Println(carts)
	return carts.TotalCost
}
func getCurrentCount(user_id string, name string) int {
	db := InitDB()
	defer db.Close()

	var carts cart.CartManagement
	db.Where("Name = ?", name).First(&carts)
	fmt.Println(carts)
	return carts.Count
}

// RemoveItemInCart creates a cart entry row in DB
func RemoveItemInCart(cart Cart) error {
	db := InitDB()
	defer db.Close()

	var fruits fruit.FruitManagement
	db.Table("fruit_managements").Where("name = ?", cart.Name).First(&fruits)

	cartId := cart.UserID + cart.UserID
	amount := fruits.Cost * getCurrentCost(cartId, *&cart.UserID, cart.Name)

	var payments payment.PaymentManagement
	db.Where("ID = ?", cartId).First(&payments)
	payment := &payment.PaymentManagement{
		ID:     &cartId,
		UserID: cart.UserID,
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
	if err != nil {
		fmt.Println("An error occurred...")
		fmt.Println(err)
	}
	return err
}

// CreateCartItem creates a cart entry row in DB
func GetPaymentAmoutFromCart(input Payment) (payment.PaymentManagement, error) {
	db := InitDB()
	defer db.Close()
	var totalAmount float64
	var cart cart.CartManagementCollection
	db.Table("cart_managements").Where("user_id = ?", input.UserID).
		Where("user_id = ?", input.UserID).Find(&cart)

	for _, x := range cart {
		totalAmount += float64(x.TotalCost)
	}
	var payments payment.PaymentManagement
	db.Model(&payments).Where("ID = ?", input.ID).
		Where("user_id = ?", input.UserID).
		Update("amount", totalAmount)

	db.Where("ID = ?", input.ID).
		Where("user_id = ?", input.UserID).
		First(&payments)

	fmt.Println("totalAmount", totalAmount)

	return payments, err
}

func PayAmount(input Payment) (payment.PaymentManagement, error) {
	db := InitDB()
	defer db.Close()
	paymentID := input.UserID + input.UserID
	var payments payment.PaymentManagement
	db.Model(&payments).
		Where("ID = ?", paymentID).
		Where("user_id = ?", input.UserID).
		Update("amount", 0).
		Update("PaymentStatus", "PAID")

	user := user.UserManagement{
		ID: input.UserID,
	}

	paymentStatus := "NOTPAID"
	paymentAmount := float64(0)
	cartItemApple := cart.CartManagement{
		UserID: user.UserID,
		Name:   "Apple",
	}
	cartItemBanana := cart.CartManagement{
		UserID: user.UserID,
		Name:   "Banana",
	}
	cartItemPear := cart.CartManagement{
		UserID: user.UserID,
		Name:   "Pear",
	}
	cartItemOrange := cart.CartManagement{
		UserID: user.UserID,
		Name:   "Orange",
	}
	payment := payment.PaymentManagement{
		ID:            &paymentID,
		UserID:        user.ID,
		PaymentStatus: &paymentStatus,
		Amount:        &paymentAmount,
	}

	err := db.Model(&cartItemApple).
		Update("Count", 0).
		Update("totalCost", 0).Error
	db.Model(&cartItemBanana).
		Update("totalCost", 0).
		Update("Count", 0)
	db.Model(&cartItemPear).
		Update("totalCost", 0).
		Update("Count", 0)
	db.Model(&cartItemOrange).
		Update("totalCost", 0).
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
