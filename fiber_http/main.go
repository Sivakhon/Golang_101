package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2" // Importing the JWT middleware for Fiber (validation jwt)
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

	app.Post("/login", login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	app.Use((checkmiddleware))

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)

	app.Post("/books", createBook)
	app.Post("/upload", uploadFile)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen(":8080")
}
