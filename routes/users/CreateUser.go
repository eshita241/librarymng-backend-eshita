package users

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateUser handles the creation of a user in the database
// Accepts a JSON object containing the user information
// Returns an HTTP status code and a message indicating the result of the operation
//
// Parameters:
//   - c: Fiber context
//
// Returns:
//   - error: An error message indicating the result of the operation and an HTTP status code
func CreateUser(c *fiber.Ctx) error {

	var user models.User
	err := c.BodyParser(&user) // Parse the request body into the User struct

	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	// Validate the user details
	/*if err := helpers.ValidateUser(user); err != nil {
		log.Printf("Error validating user: %v\n", err)
		return c.Status(400).SendString("Error validating user")
	}*/

	// Ensure a role is provided
	if user.Role != "librarian" && user.Role != "staff" && user.Role != "member" {
		log.Printf("Invalid user role: %v\n", user.Role)
		return c.Status(400).SendString("Invalid user role")
	}

	// Create the user in the database
	result := database.Database.Db.Create(&user)

	if result.Error != nil {
		log.Printf("Error creating user: %v\n", result.Error)
		return c.Status(500).SendString("Error creating user")
	}

	log.Printf("User with ID %v created\n", user.ID)
	return c.Status(200).JSON(user)
}
