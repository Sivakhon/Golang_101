package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&Book{}, &User{})

	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		return registerUser(db, c)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		token, err := loginUser(db, c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	app.Use("/books", authRequired)

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
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		return updateBook(db, c)
	})
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		return deleteBook(db, c)
	})

	// Start server
	log.Fatal(app.Listen(":8000"))
}
