package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)



func getBooks(c fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c fiber.Ctx) error {
	bookIdStr := c.Params("id")
	boolId, err := strconv.Atoi(bookIdStr)
	if err != nil || boolId < 0 || boolId >= len(books) {
		return c.Status(400).SendString("Invalid book ID")
	}

	return c.JSON(books[int(boolId)])
}

func createbook(c fiber.Ctx) error {
	book := new(Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	books = append(books, *book)
	return c.Status(200).SendString("Book created successfully")
}

func updateBook(c fiber.Ctx) error {
	bookIdStr := c.Params("id")

	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil || bookId < 0 || bookId >= len(books) {
		return c.Status(400).SendString("Invalid book ID")
	}

	bookUpdate := new(Book)
	if err := c.Bind().Body(bookUpdate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	books[bookId].Title = bookUpdate.Title
	books[bookId].Author = bookUpdate.Author

	return c.Status(200).SendString("Book updated successfully")
}

func deleteBook(c fiber.Ctx) error {
	bookIdStr := c.Params("id")

	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil || bookId < 0 || bookId >= len(books) {
		return c.Status(400).SendString("Invalid book ID")
	}

	books = append(books[:bookId], books[bookId+1:]...)

	return c.Status(200).SendString("Book delete successfully")
}

func uploadFile(c fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = c.SaveFile(file, "./upload/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully")
}

func loginBook(c fiber.Ctx) error {
	user := new(User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if user.Username != memberuser.Username || user.Password != memberuser.Password {
		return fiber.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["Username"] = user.Username
	claims["role"] = "admin"
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token
    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
    	return c.SendStatus(fiber.StatusInternalServerError)
    }

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   t,
	})
}
