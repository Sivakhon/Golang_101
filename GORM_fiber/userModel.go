package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func registerUser(db *gorm.DB, c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "This email is already registered",
		})
	}
	return nil
}

func loginUser(db *gorm.DB, c *fiber.Ctx) (string, error) {

	// Parse the request body to get the login credentials
	login_user := new(User)
	if err := c.BodyParser(login_user); err != nil {
		return "", err
	}

	exist_user := new(User)
	result := db.Where("email = ?", login_user.Email).First(&exist_user)
	if result.Error != nil {
		return "", result.Error
	}
	// Compare password: hash in DB vs plain password input
	err := bcrypt.CompareHashAndPassword([]byte(exist_user.Password), []byte(login_user.Password))
	if err != nil {
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = login_user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return t, nil
}
