package main

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint
	Name        string
	Author      string
	Description string
	Price       int
}

// func CreateBook(db *gorm.DB, book *Book) {
// 	result := db.Create(book)
// 	if result.Error != nil {
// 		log.Fatalf("Error creating book: %v", result.Error)
// 	}
// 	fmt.Println("Book created successfully")
// }

func GetBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error finding book: %v", result.Error)
	}
	return &book
}

// func UpdateBook(db *gorm.DB, book *Book) {
// 	result := db.Save(book)
// 	if result.Error != nil {
// 		log.Fatalf("Error updating book: %v", result.Error)
// 	}
// 	fmt.Println("Book updated successfully")
// }

// func DeleteBook(db *gorm.DB, id uint) {
// 	var book Book
// 	result := db.Unscoped().Delete(&book, id)
// 	if result.Error != nil {
// 		log.Fatalf("Error deleting book: %v", result.Error)
// 	}
// 	fmt.Println("Book deleted successfully")
// }

func searchBook(db *gorm.DB, bookName string) []Book {
	var books []Book
	result := db.Where("name = ?", bookName).Find(&books)

	if result.Error != nil {
		log.Fatalf("Error deleting book: %v", result.Error)
	}
	return books
}
