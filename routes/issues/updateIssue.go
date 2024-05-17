package issues

// Due dates and notes

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateIssue(c *fiber.Ctx) error {
	//create a variable for Book struct
	var updatedIssue models.Issue

	// parse the body and error handling
	err := c.BodyParser(&updatedIssue)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).JSON(err.Error())
	}

	// validates book, display error if not validated
	if err := validator.ValidateIssue(updatedIssue); err != nil {
		log.Printf("Error validating issue: %v\n", err)
		return c.Status(400).SendString("Error validating Issue")
	}

	// update book
	result := database.Database.Db.Model(&models.Issue{}).Where("id = ?", updatedIssue.ID).Updates(&updatedIssue)

	// error handling
	if result.Error != nil {
		log.Printf("Error updating Issue: %v\n", result.Error)
		return c.Status(500).JSON(err.Error())
	}

	if result.RowsAffected == 0 {
		log.Printf("Issue with ID %v not found\n", updatedIssue.ID)
		return c.Status(500).JSON(err.Error())
	}

	// Success
	log.Printf("Issue with ID %v has been updated", updatedIssue.ID)
	return c.Status(200).JSON(updatedIssue)
}
