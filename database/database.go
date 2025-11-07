package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.db"))

	if err != nil {
		fmt.Println(err)
	}

	DB = db

	fmt.Println("Connected to database")
}
