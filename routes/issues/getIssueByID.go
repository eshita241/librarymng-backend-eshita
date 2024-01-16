package issues

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

func GetIssue(c *fiber.Ctx) error {

	// Get issue ID from URL parameters
	issueID, err := strconv.Atoi(c.Params("id"))

	// error handling
	if err != nil {
		log.Printf("Error converting issue ID to integer: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid issue ID"})
	}

	// Create a variable issue
	var issue models.Issue

	// validates issue, display error if not validated
	if err := helpers.ValidateIssue(issue); err != nil {
		log.Printf("Error validating issue: %v\n", err)
		return c.Status(400).SendString("Error validating issue")
	}

	// Fetch the issue from the database
	result := database.Database.Db.First(&issue, issueID)

	// Handle errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Issue with ID %v not found\n", issueID)
			return c.Status(404).JSON(fiber.Map{"error": "Issue not found"})
		}
		log.Printf("Error fetching issue: %v\n", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Success
	log.Printf("Issue with ID %v has been fetched\n", issueID)
	return c.Status(200).JSON(issue)
}
