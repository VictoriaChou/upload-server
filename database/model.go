package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User response to user table
type User struct {
	gorm.Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255);unique_index:unique_user"` // create unique index for user called unique user
	Age  int16
}

// Get get first user satisfy the condition
func (u *User) Get(query *gorm.DB) error {
	err := query.Where(u).First(u).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
