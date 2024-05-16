package main //In Go, a package is a way to organize and reuse code. It is a collection of Go source files that are organized together in a directory.

import (
	"librarymng-backend/database"
	"librarymng-backend/initializers"
	"librarymng-backend/routes/books"
	"librarymng-backend/routes/issues"
	"librarymng-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// User Routes
	app.Post("/users", users.CreateUser)
	app.Delete("/usersdel/:id", users.DeleteUser)
	app.Get("/userget/:id", users.GetUser)
	app.Put("/userupdate/:id", users.UpdateUser)

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

func init() { //connect database
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	database.ConnectToDB(&config)
}
func main() {
	app := fiber.New() //The line app := fiber.New() creates a new instance of a Fiber application.
	//Fiber is a web framework for Go (Golang) that is designed to be fast, flexible, and easy to use. It is built on top of Fasthttp, which is a high-performance HTTP server implementation for Go.
	app.Use(logger.New()) //app.Use() attaches middleware to middleware stack; logger.New sets up logger middlware

	SetupRoutes(app) //connect all the routes of app

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error { //basic router for server side connection check
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000")) //running on port 8000
}
