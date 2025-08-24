package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	// newBook := Book{
	// 	Name:        "The Go Programming Language",
	// 	Author:      "Alan A. A. Donovan",
	// 	Description: "Comprehensive guide to Go",
	// 	Price:       250}
	// // CreateBook(db, &newBook)

	// // Get a book
	book := GetBook(db, 3) // Assuming the ID of the book is 1
	fmt.Println("Book Retrieved:", book)

	// // Update a book
	// book.Name = "The Go , Updated Edition"
	// book.Price = 2545
	// UpdateBook(db, book)

	// Delete a book
	// DeleteBook(db, 1)

	//search
	currentBook := searchBook(db, "The Go Programming Language") // Assuming the ID of the book is 1
	for _, book := range currentBook {
		fmt.Println(book.ID, book.Name, book.Description, book.Author, book.Price)
	}
}
