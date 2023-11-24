package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web")

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendFile("./web/about.html")
	})
	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.SendFile("./web/projects.html")
	})
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendFile("./web/contact.html")
	})

	app.Listen(":3000")
}
