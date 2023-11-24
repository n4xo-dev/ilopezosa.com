package main

import (
	"fmt"
	"strings"

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

	app.Post("/contact", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		message := c.FormValue("message")

		// Send email
		// ...
		fmt.Println(name, email, message)
		if strings.ToLower(name) == "error" {
			return c.SendFile("./web/cards/error.html")
		}
		return c.SendFile("./web/cards/success.html")
	})

	app.Listen(":3000")
}
