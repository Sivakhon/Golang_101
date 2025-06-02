package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func checkmiddleware(c fiber.Ctx) error {

	start := time.Now()

	fmt.Printf("Url = %s, Method = %s, Time =  %s\n", c.Path(), c.Method(), start)

	return c.Next()
}
