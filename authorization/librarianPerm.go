package authorization

import (
	"librarymng-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthLibrarian is a middleware to check if the user has the librarian role
func AuthLibrarian(c *fiber.Ctx) error {
	// This example assumes the role is stored in a header for simplicity
	// Adjust this part to fit your authentication method (JWT, session, etc.)
	user := c.Locals("user").(models.AuthResponse)

	if user.Role != "librarian" {
		log.Printf("Unauthorized access attempt: Role is %v\n", user.Role)
		return c.Status(403).SendString("Unauthorized")
	}

	return c.Next()
}
