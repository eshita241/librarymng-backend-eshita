package books

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateBook(c *fiber.Ctx) error {
	//create a variable for Book struct
	var updatedBook models.Book

	// parse the body and error handling
	err := c.BodyParser(&updatedBook)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).JSON(err.Error())
	}

	// validates book, display error if not validated
	if err := validator.ValidateBook(updatedBook); err != nil {
		log.Printf("Error validating book: %v\n", err)
		return c.Status(400).SendString("Error validating book")
	}

	// update book
	result := database.Database.Db.Model(&models.Book{}).Where("id = ?", updatedBook.ID).Updates(&updatedBook)

	// error handling
	if result.Error != nil {
		log.Printf("Error updating book: %v\n", result.Error)
		return c.Status(500).JSON(err.Error())
	}

	if result.RowsAffected == 0 {
		log.Printf("Book with ID %v not found\n", updatedBook.ID)
		return c.Status(500).JSON(err.Error())
	}

	// Success
	log.Printf("Book with ID %v has been updated", updatedBook.ID)
	return c.Status(200).JSON(updatedBook)
}
