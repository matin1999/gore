package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Scheduler service is running!")
	})

	go func() {
		
	}()

	log.Fatal(app.Listen(":3000"))
}
