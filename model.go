package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func createUser(db *gorm.DB, c *fiber.Ctx) error {

	//สร้างuser เป็นpointerไปยัง struct Userเพื่อเอาไว้เก็บข้อมูลที่รับมาจาก request
	user := new(User)

	//BodyParser ของ Fiber จะอ่าน JSON body
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//การhash เข้ารหัสโดยbcryptสร้างเป็นbyte table user colum password เข้ารหัสและทำการวน
	//bcrypt.DefaultCost คือ cost , int โดยปกติintจะเป็น 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	//แปลง hashedPassword เป็น string
	user.Password = string(hashedPassword)

	//ทำการเก็บ password ที่เข้ารหัสแล้วลงใน user
	db.Create(&user)

	return c.JSON(user)

}

// get all books
func GetBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Fatalf("Error finding book: %v", result.Error)
	}
	return books
}

func CreateBook(db *gorm.DB, book *Book) {
	result := db.Create(book)
	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}
	fmt.Println("Book created successfully")
}

func GetBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error finding book: %v", result.Error)
	}
	return &book
}

func UpdateBook(db *gorm.DB, book *Book) error {
	result := db.Save(book)
	if result.Error != nil {
		return fmt.Errorf("Error updating book: %v", result.Error)
	}
	return nil
}

func DeleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Unscoped().Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Error deleting book: %v", result.Error)
	}
	fmt.Println("Book deleted successfully")
}
