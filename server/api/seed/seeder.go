package seed

import (
	"fruitshop/server/api/models"
	"log"

	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&Fruit{},
		&SingleItemDiscount{},
		&DualItemDiscount{},
		&SingleItemCoupon{},
		&AppliedDualItemDiscount{},
		&AppliedSingleItemCoupon{},
		&AppliedSingleItemDiscount{},
		&Payment{},
		&Cart{},
		&CartItem{},
		&Customer{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&Fruit{},
		&SingleItemDiscount{},
		&DualItemDiscount{},
		&SingleItemCoupon{},
		&AppliedDualItemDiscount{},
		&AppliedSingleItemCoupon{},
		&AppliedSingleItemDiscount{},
		&Payment{},
		&Cart{},
		&CartItem{},
		&Customer{}).Error
	if err != nil {
		log.Fatalf("Unable to migrate tables: %v", err)
	}

	loadFruitsAndDiscountsTableMetaData()

}

func loadFruitsAndDiscountsTableMetaData() {

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

}
