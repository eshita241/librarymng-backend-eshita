package issues

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetIssueHistory handles the GET request to retrieve the issue history for a user by user ID
func GetIssueHistory(c *fiber.Ctx) error {
	// Get the user ID from the request parameters
	userID, err := strconv.Atoi(c.Params("id"))

	// error handling
	if err != nil {
		log.Printf("Error converting user ID to integer: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Query the database to get the issue history for the specified user ID
	var issues []models.Issue
	result := database.Database.Db.Preload("Book").Preload("User").Where("user_id = ?", userID).Find(&issues)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("No issues found for User with ID %v\n", userID)
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found or no issues found"})
		}
		log.Printf("Error fetching issue history for User with ID %v: %v\n", userID, result.Error)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch issue history"})
	}

	// Return the issue history as JSON
	log.Printf("All Issues of User with ID %v has been fetched\n", userID)
	return c.Status(http.StatusOK).JSON(issues)
}
