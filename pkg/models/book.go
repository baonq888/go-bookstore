package models

import (
	"go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) 
}

func (b *Book) CreateBook() *Book {
	query := "INSERT INTO books (name, author, publication, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW()) RETURNING id"
	err := db.Raw(query, b.Name, b.Author, b.Publication).Scan(&b.ID).Error
	if err != nil {
		return nil
	}
	return b
}

func GetAllBooks() []Book {
	var books []Book
	query := "SELECT id, name, author, publication, created_at, updated_at FROM books"
	db.Raw(query).Scan(&books)
	return books
}

func GetBookById(Id int64) (*Book, error) {
	var book Book
	query := "SELECT id, name, author, publication, created_at, updated_at FROM books WHERE id = ?"
	err := db.Raw(query, Id).Scan(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) UpdateBook(Id int64) error {
	query := "UPDATE books SET name = ?, author = ?, publication = ?, updated_at = NOW() WHERE id = ?"
	result := db.Exec(query, b.Name, b.Author, b.Publication, Id)
	return result.Error
}

func DeleteBook(Id int64) error {
	query := "DELETE FROM books WHERE id = ?"
	result := db.Exec(query, Id)
	return result.Error
}