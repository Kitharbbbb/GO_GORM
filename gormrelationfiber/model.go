package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
}

func createPublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Create(publisher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createBookWithAuthor(db *gorm.DB, book *Book, authorIDs []uint) error {
	// First, create the book
	if err := db.Create(book).Error; err != nil {
		return err
	}

	// add authors (many2many)
	if len(authorIDs) > 0 {
		var authors []Author
		if err := db.Where("id IN ?", authorIDs).Find(&authors).Error; err != nil {
			return err
		}
		if err := db.Model(book).Association("Authors").Append(&authors); err != nil {
			return err
		}
	}
	return nil
}

func getBookWithPublisher(db *gorm.DB, bookID uint, c *fiber.Ctx) error {
	id := c.Query("id")
	var book Book
	result := db.Preload("Publisher").First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func getBookWithAuthors(db *gorm.DB, bookID uint, c *fiber.Ctx) error {
	id := c.Query("id")
	var book Book
	result := db.Preload("Authors").First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func listBooksOfAuthor(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Query("author_id")
	var books []Book
	result := db.Joins("JOIN author_books on author_books.book_id = books.id").
		Where("author_books.author_id = ?", id).
		Find(&books)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(books)
}
