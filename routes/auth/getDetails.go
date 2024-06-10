package auth

import (
	"librarymng-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	// Retrieve authenticated user from Fiber's locals
	user := c.Locals("user").(models.AuthResponse) // Assuming your user model is models.Auth

	// Return the user as JSON response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}
