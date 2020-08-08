package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllFruits(t *testing.T) {

	err := refreshFruitTable()
	if err != nil {
		log.Fatalf("Error refreshing fruits table %v\n", err)
	}

	err = seedFruits()
	if err != nil {
		log.Fatalf("Error seedFruits %v\n", err)
	}

	users, err := fruitInstance.FindAllFruits(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the fruits: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 4)
}
