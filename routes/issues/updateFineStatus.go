package issues

import (
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateFineStatus(c *fiber.Ctx) error {
	// Get issue ID from URL parameters
	issueID, err := strconv.Atoi(c.Params("id"))

	// error handling
	if err != nil {
		log.Printf("Error converting issue ID to integer: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid issue ID"})
	}

	// Create a variable issue
	var issue models.Issue

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

	// calculate days late
	daysLate := int(issue.ReturnDate.Sub(issue.DueDate).Hours() / 24)

	// Check if the book is returned late
	if daysLate > 0 {
		// Update the fine (you can adjust the fine calculation as needed)
		issue.Fine = uint(daysLate * 100) // Assuming 100 rupees per day as an example
		// Save the updated issue back to the database
		if err := database.Database.Db.Save(&issue).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update fine"})
		}
		return c.JSON(fiber.Map{"message": "Fine updated successfully"})
	}

	// Status Update
	issue.Status = "returned"
	// Save the updated issue back to the database
	if err := database.Database.Db.Save(&issue).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update status"})
	}

	return c.JSON(fiber.Map{"message": "Book returned on time, no fine applied"})
}
