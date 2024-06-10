package main //In Go, a package is a way to organize and reuse code. It is a collection of Go source files that are organized together in a directory.

import (
	"librarymng-backend/authorization"
	"librarymng-backend/database"
	"librarymng-backend/initializers"
	"librarymng-backend/middleware"
	"librarymng-backend/routes/auth"
	"librarymng-backend/routes/books"
	"librarymng-backend/routes/issues"
	"librarymng-backend/routes/users"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/patrickmn/go-cache"
)

var cacheInstance *cache.Cache

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

	cacheInstance = cache.New(10*time.Minute, 20*time.Minute)

	//cache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time.

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(middleware.CacheMiddleware(cacheInstance))

	SetupRoutes(app) //connect all the routes of app

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error { //basic router for server side connection check
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000")) //running on port 8000
}

func SetupRoutes(app *fiber.App) {
	//Auth Routes
	app.Route("/auth", func(router fiber.Router) {
		router.Post("/register", auth.SignUpUser)
		router.Post("/login", auth.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, auth.LogoutUser)
	})
	app.Get("/users/me", middleware.DeserializeUser, auth.GetMe)

	// User Routes
	app.Post("/users", users.CreateUser)
	app.Delete("/usersdel/:id", users.DeleteUser)
	app.Get("/userget/:id", users.GetUser)
	app.Put("/userupdate/:id", users.UpdateUser)

	// Book Routes
	app.Post("/api/book", middleware.DeserializeUser, authorization.AuthLibrarian, books.AddBook)
	app.Put("/api/bookup/:id", middleware.DeserializeUser, authorization.AuthLibrarian, books.UpdateBook)
	app.Delete("/api/bookdel/:id", middleware.DeserializeUser, authorization.AuthLibrarian, books.DeleteBook)
	app.Get("/api/bookget/:id", func(c *fiber.Ctx) error {
		return books.GetBook(c, cacheInstance)
	})
	app.Get("/api/books/search", middleware.DeserializeUser, authorization.AuthLibrarian, books.SearchBooks)

	// Issue Routes
	app.Post("/api/issue", middleware.DeserializeUser, authorization.AuthMember, issues.AddIssue)
	app.Get("/api/issuegethis/:id", middleware.DeserializeUser, authorization.AuthMember, issues.GetIssueHistory)
	app.Get("/api/issuegetid/:id", middleware.DeserializeUser, authorization.AuthLibrarian, issues.GetIssue)
	app.Put("/api/issueup/:id", middleware.DeserializeUser, authorization.AuthMember, issues.UpdateIssue)
	app.Put("/api/issueupfs/:id", middleware.DeserializeUser, authorization.AuthLibrarian, issues.UpdateFineStatus)
}

/*
LF (Line Feed):
LF is a control character used to start a new line in text files. It is commonly used on Unix and Unix-like systems (including Linux and macOS) as the standard line ending character.
CRLF (Carriage Return + Line Feed):
CRLF consists of two control characters: carriage return followed by line feed. It is the standard line ending sequence used in text files on Windows systems.


Middleware is a software architectural pattern that allows you to encapsulate and separate concerns in your application by adding layers of functionality to the request-response cycle.
In web development, middleware functions are used to process HTTP requests and responses before they reach the final handler.
Modularity, Chain of Responsibility, Request Processing


Semantic Versioning (SemVer):
Semantic versioning is a versioning scheme that specifies how version numbers are assigned and incremented for software releases. It follows the format MAJOR.MINOR.PATCH, where:
MAJOR is incremented for incompatible API changes,
MINOR is incremented for backward-compatible feature additions, and
PATCH is incremented for backward-compatible bug fixes.
++incompatible Warning:
The ++incompatible notation indicates that the dependency does not adhere to semantic versioning conventions.
This warning is shown when the Go module system detects that the versioning scheme of a dependency is incompatible with the expected semantic versioning format.
*/
