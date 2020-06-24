package contextdb

import (
	"github.com/gecapo/rest-api-golang/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//DB pointer to the database context
var DB *gorm.DB

//ConnectDataBase opens connetion to db to db
func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.RecipeDBO{})

	DB = database
}
