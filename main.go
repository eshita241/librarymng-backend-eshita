package main

import (
	"librarymng-backend/database"

	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

}

func main() {
	app := fiber.New()

	database.ConnectToDB()

	SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	log.Fatal(app.Listen(":3000"))
}
