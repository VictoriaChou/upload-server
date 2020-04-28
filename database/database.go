package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // pg driver
)

type User struct {
	gorm.Model
	Name string
	Age  int16
}

func main() error {
	db, err := gorm.Open("172.17.0.2", "postgres")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Scarlett", Age: 18})

	// Read
	var user User
	db.First(&user, "name = ?", "Scarlett")
}
