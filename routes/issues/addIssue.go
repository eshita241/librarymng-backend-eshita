package issues

import (
	"librarymng-backend/database"
	"librarymng-backend/helpers"
	"librarymng-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AddIssue(c *fiber.Ctx) error {
	//create a variable for Book struct
	var issue models.Issue

	// parse the body and error handling
	err := c.BodyParser(&issue)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).JSON(err.Error())
	}

	// validates issue, display error if not validated
	if err := helpers.ValidateIssue(issue); err != nil {
		log.Printf("Error validating issue: %v\n", err)
		return c.Status(400).SendString("Error validating issue")
	}

	// save the book in database
	result := database.Database.Db.Create(&issue)

	// error handling
	if result.Error != nil {
		log.Printf("Error creating booking: %v\n", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Database error", "details": result.Error.Error()})
	}

	// return the book in JSON format
	return c.Status(200).JSON(fiber.Map{
		"message": "Issue created successfully",
	})
}
