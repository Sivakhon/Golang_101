package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func checkmiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("Url = %s, Method = %s, Time =  %s\n", c.Path(), c.Method(), start)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"] != true {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Username == memberuser.Username && user.Password == memberuser.Password {

		// Create the Claims
		claims := jwt.MapClaims{
			"username": user.Username,
			"admin":    true,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"message": "Login successful",
			"user":    user.Username,
			"token":   t,
		})
	} else {
		return fiber.ErrUnauthorized
	}

}
