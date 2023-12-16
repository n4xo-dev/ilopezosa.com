package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web")

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendFile("./web/home.html")
	})
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
		time.Sleep(5 * time.Second)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
