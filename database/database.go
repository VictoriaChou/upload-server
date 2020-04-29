package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // pg driver
)

var MyDB *gorm.DB

func Init() {
	host := "127.0.0.1"
	port := "5432"
	dbUser := "postgres"
	dbName := "postgres"
	password := "secret"

	var err error
	connCfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, dbUser, dbName, password)

	MyDB, err = gorm.Open("postgres", connCfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
