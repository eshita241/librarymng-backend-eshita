package books

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	//create a variable for Book struct
	var book models.Book

	// parse the body and error handling
	err := c.BodyParser(&book)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).JSON(err.Error())
	}

	// validates book, display error if not validated
	if err := validator.ValidateBook(book); err != nil {
		log.Printf("Error validating book: %v\n", err)
		return c.Status(400).SendString("Error validating book")
	}

	// save the book in database
	result := database.Database.Db.Create(&book)

	// error handling
	if result.Error != nil {
		log.Printf("Error creating booking: %v\n", result.Error)
		return c.Status(500).JSON(err.Error())
	}

	// return the book in JSON format
	return c.Status(200).JSON(fiber.Map{
		"message": "Book added successfully",
		"book":    book,
	})
}
