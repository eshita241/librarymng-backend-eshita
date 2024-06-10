package authorization

import (
	"librarymng-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthLibrarian is a middleware to check if the user has the librarian role
func AuthMember(c *fiber.Ctx) error {
	// This example assumes the role is stored in a header for simplicity
	// Adjust this part to fit your authentication method (JWT, session, etc.)
	response := c.Locals("user").(models.AuthResponse)

	if response.Role != "member" {
		log.Printf("Unauthorized access attempt: Role is %v\n", response.Role)
		return c.Status(403).SendString("Unauthorized")
	}

	return c.Next()
}
