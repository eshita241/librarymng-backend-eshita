package main

import (
	"librarymng-backend/database"
	"librarymng-backend/routes/books"

	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// User Routes
	// Book Routes
	app.Post("/api/book", books.AddBook)
	app.Put("/api/bookup/:id", books.UpdateBook)
	app.Delete("/api/bookdel/:id", books.DeleteBook)
	app.Get("/api/bookget/:id", books.GetBook)
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
