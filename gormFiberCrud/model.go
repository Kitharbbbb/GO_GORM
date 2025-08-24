package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /book [get]

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// getBooks retrieves all books
func getBooks(db *gorm.DB, c *fiber.Ctx) error {
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

// getBook retrieves a book by id
func getBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	db.First(&book, id)
	return c.JSON(book)
}

// createBooks create new book
func createBook(db *gorm.DB, c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	db.Create(&book)
	return c.JSON(book)
}

// updateBook
func updateBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)
	db.First(&book, id)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	db.Save(&book)
	return c.JSON(book)
}

// deleteBook
func deleteBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	db.Delete(&Book{}, id)
	return c.SendString("success delete")
}
