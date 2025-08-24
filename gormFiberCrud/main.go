package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8000
// @BasePath /
// @schemes http
// @in header
// @name Authorization

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

	// Migrate the schema
	db.AutoMigrate(&Book{})

	// Setup Fiber
	app := fiber.New()

	// CRUD routes
	app.Get("/books", func(c *fiber.Ctx) error {
		return getBooks(db, c)
	})
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		return getBook(db, c)
	})
	app.Post("/books", func(c *fiber.Ctx) error {
		return createBook(db, c)
	})
	app.Put("/book/:id", func(c *fiber.Ctx) error {
		return updateBook(db, c)
	})
	app.Delete("/book/:id", func(c *fiber.Ctx) error {
		return deleteBook(db, c)
	})
	// app.Put("/books/:id", func(c *fiber.Ctx) error {
	// 	return updateBook(db, c)
	// })
	// app.Delete("/books/:id", func(c *fiber.Ctx) error {
	// 	return deleteBook(db, c)
	// })

	// Start server
	log.Fatal(app.Listen(":8000"))
}
