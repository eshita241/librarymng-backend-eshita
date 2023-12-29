package main

import (
	"librarymng-backend/database"
	"librarymng-backend/routes/books"
	"librarymng-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/book", books.AddBook)
	app.Put("/api/bookup/:id", books.UpdateBook)
	app.Delete("/api/bookdel/:id", books.DeleteBook)
	app.Get("/api/bookget/:id", books.GetBook)
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
