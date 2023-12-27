package books

import (
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/helpers"
	"librarymng-backend/models"
	"log"
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
	if err := helpers.ValidateBook(book); err != nil {
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
}
