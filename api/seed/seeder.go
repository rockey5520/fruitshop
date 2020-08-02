package seed

import (
	"log"

	"fruitshop/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
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

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(
		&models.Post{},
		&models.User{},
		&models.Fruit{},
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.Payment{},
		&models.Cart{},
		&models.CartItem{},
	).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(
		&models.Post{},
		&models.User{},
		&models.Fruit{},
		&models.SingleItemDiscount{},
		&models.DualItemDiscount{},
		&models.SingleItemCoupon{},
		&models.Payment{},
		&models.Cart{},
		&models.CartItem{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	appleItemDiscount := models.SingleItemDiscount{Count: 7, Discount: 10, Name: "APPLE10"}
	orangeSingleItemCoupon := models.SingleItemCoupon{
		Discount: 30,
		Name:     "ORANGE30",
	}
	apple := models.Fruit{
		Name: "Apple",
		SingleItemDiscount: []models.SingleItemDiscount{
			appleItemDiscount,
		},
		Price: 1,
	}
	banana := models.Fruit{
		Name:  "Banana",
		Price: 1,
	}
	pear := models.Fruit{
		Name:  "Pear",
		Price: 1,
	}
	orange := models.Fruit{
		Name: "Orange",
		SingleItemCoupon: []models.SingleItemCoupon{
			orangeSingleItemCoupon,
		},
		Price: 1,
	}

	if err := db.Create(&apple).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&banana).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&pear).Error; err != nil {
		panic("Unable to create fruit inventory")
	}
	if err := db.Create(&orange).Error; err != nil {
		panic("Unable to create fruit inventory")
	}

	var pearFromDB models.Fruit
	db.Where("name = ?", "Pear").First(&pearFromDB)
	var bananaFromDB models.Fruit
	db.Where("name = ?", "Banana").First(&bananaFromDB)
	dualItemDiscount := models.DualItemDiscount{
		FruitID:   pearFromDB.ID,
		FruitID_1: pearFromDB.ID,
		FruitID_2: bananaFromDB.ID,
		Count_1:   4,
		Count_2:   2,
		Name:      "PEARBANANA30",
		Discount:  30,
	}
	if err := db.Create(&dualItemDiscount).Error; err != nil {
		panic("Unable to create Single item discount inventory")
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
