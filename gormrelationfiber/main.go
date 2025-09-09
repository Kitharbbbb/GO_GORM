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

	db.AutoMigrate(&Book{}, &Publisher{}, &Author{}, &AuthorBook{})

	// ขาสร้าง
	publisher := Publisher{
		Details: "Publisher Details",
		Name:    "Publisher Name",
	}
	_ = createPublisher(db, &publisher)

	// Example data for a new author
	author := Author{
		Name: "Author Name",
	}
	_ = createAuthor(db, &author)

	// // Example data for a new book with an author
	book := Book{
		Name:        "Book Title",
		Author:      "Book Author",
		Description: "Book Description",
		PublisherID: publisher.ID,     // Use the ID of the publisher created above
		Authors:     []Author{author}, // Add the created author
	}
	_ = createBookWithAuthor(db, &book, []uint{author.ID})

	app := fiber.New()
	app.Get("/BookWithPublisher", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id") // ดึงจาก path param
		return getBookWithPublisher(db, uint(id), c)
	})
	app.Get("/BookWithAuthors", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id") // ดึงจาก path param
		return getBookWithAuthors(db, uint(id), c)
	})
	app.Get("/listBooksOfAuthor", func(c *fiber.Ctx) error {
		return listBooksOfAuthor(db, c)
	})
	log.Fatal(app.Listen(":3000"))

}
