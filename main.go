package main

import (
	"librarymng-backend/database"
	"librarymng-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", users.CreateUser)
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
