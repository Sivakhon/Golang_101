package main

import (
    "os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var books []Book
var memberuser = User{
	Username: "admin@gmail.com",
	Password: "password1234",
}

func main() {

	books = append(books, Book{ID: 0, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 1, Title: "Brave New World", Author: "Aldous Huxley"})
	books = append(books, Book{ID: 2, Title: "Fahrenheit 451", Author: "Ray Bradbury"})

	app := fiber.New()

	app.Post("/login", loginBook)

	app.Use((checkmiddleware))

	app.Use(jwtware.New(jwtware.Config{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/books", createbook)
	app.Post("/upload", uploadFile)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen(":8080")
}