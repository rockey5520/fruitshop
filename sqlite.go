package users
import (
        	"fruitshop/gen/user"
        	"github.com/jinzhu/gorm"
        	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
var db *gorm.DB
var err error
type User *user.UserManagement
 
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
        	// Creating the table if it doesn't exist
        	var TableStruct = user.UserManagement{}
        	if !db.HasTable(TableStruct) {
                    	db.CreateTable(TableStruct)
                    	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(TableStruct)
        	}
 
        	return db
}
 
// GetClient retrieves one client by its ID
func GetUser(UserEmailID string) (user.UserManagement, error) {
        	db := InitDB()
        	defer db.Close()
 
        	var users user.UserManagement
        	db.Where("UserEmailID = ?", UserEmailID).First(&users)
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