package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Atanu@616@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}else {
		fmt.Println("Database connected successfully")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
