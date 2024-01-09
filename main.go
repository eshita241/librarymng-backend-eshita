package main

// postgres://postgres:Itba@1234..78@localhost:5432/librarymng?sslmode=disable
import (
	"librarymng-backend/database"
	"librarymng-backend/routes/books"
	"librarymng-backend/routes/issues"
	"librarymng-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// User Routes
	app.Post("/users", users.CreateUser)

	// Book Routes
	app.Post("/api/book", books.AddBook)
	app.Put("/api/bookup/:id", books.UpdateBook)
	app.Delete("/api/bookdel/:id", books.DeleteBook)
	app.Get("/api/bookget/:id", books.GetBook)

	// Issue Routes
	app.Post("/api/issue", issues.AddIssue)
	app.Get("/api/issuegethis/:id", issues.GetIssueHistory)
	app.Get("/api/issuegetid/:id", issues.GetIssue)
	app.Put("/api/issueup/:id", issues.UpdateIssue)
	app.Put("/api/issueupfs/:id", issues.UpdateFineStatus)
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
