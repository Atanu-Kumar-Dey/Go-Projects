package models

import (
	"book-store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Books struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Books{})
}

func (b *Books) CreateBook() *Books {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Books {
	var Books []Books
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Books, *gorm.DB) {
	var getBooks Books
	db := db.Where("Id=?", Id).Find(&getBooks)
	return &getBooks, db
}

func DeleteBook(Id int64) Books {
	var deleteBook Books
	db.Where("Id=?", Id).Delete(deleteBook)
	return deleteBook
}
