/*package books

import (
	"encoding/json"
	"io/ioutil"
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBook(c *fiber.Ctx) error {

	// Get book ID from URL parameters
	bookID, err := strconv.Atoi(c.Params("id"))

	// error handling
	if err != nil {
		log.Printf("Error converting book ID to integer: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	// Create a variable book
	var book models.Book

	// validates book, display error if not validated
	if err := validator.ValidateBook(book); err != nil {
		log.Printf("Error validating book: %v\n", err)
		return c.Status(400).SendString("Error validating book")
	}

	// Fetch the book from the database
	result := database.Database.Db.First(&book, bookID)

	// Handle errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Book with ID %v not found\n", bookID)
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}
		log.Printf("Error fetching book: %v\n", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Success
	log.Printf("Book with ID %v has been fetched\n", bookID)
	return c.Status(200).JSON(book)
}*/

package books

import (
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

func GetBook(c *fiber.Ctx, cacheInstance *cache.Cache) error {
	// Get book ID from URL parameters
	bookID, err := strconv.Atoi(c.Params("id"))

	// Error handling
	if err != nil {
		log.Printf("Error converting book ID to integer: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	cacheKey := c.Path() + "?" + c.Params("id")

	// Check cache first
	if cached, found := cacheInstance.Get(cacheKey); found {
		log.Printf("Book with ID %v found in cache\n", bookID)
		return c.Status(200).JSON(cached)
	}

	// Create a variable book
	var book models.Book

	// Fetch the book from the database
	result := database.Database.Db.First(&book, bookID)

	// Handle errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Book with ID %v not found\n", bookID)
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}
		log.Printf("Error fetching book: %v\n", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Validate the book
	if err := validator.ValidateBook(book); err != nil {
		log.Printf("Error validating book: %v\n", err)
		return c.Status(400).SendString("Error validating book")
	}

	// Cache the book data
	cacheInstance.Set(cacheKey, book, 10*time.Minute)

	// Success
	log.Printf("Book with ID %v has been fetched and cached\n", bookID)
	return c.Status(200).JSON(book)
}
