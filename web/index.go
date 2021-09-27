package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
	})
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Use(func(c *fiber.Ctx) error {
		t := time.Now()
		fmt.Println(t.Day())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/greet/:name", func(c *fiber.Ctx) error {
		if c.Params("name") == "sarthak" {
			return fiber.NewError(782, "Can't greet")
		}
		return c.SendString("Hello," + c.Params("name"))
	})

	api := app.Group("/api/", func(c *fiber.Ctx) error {
		return c.SendString("inside api")
	})

	api.Get("/register", func(c *fiber.Ctx) error { return c.SendString("Inside register") })

	app.Listen(":3000")
}
