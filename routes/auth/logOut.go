package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// LogoutUser is a handler function to log out a user by expiring their authentication token.
func LogoutUser(c *fiber.Ctx) error {
	// Set the expiration time to 24 hours ago to invalidate the cookie.
	expired := time.Now().Add(-time.Hour * 24)

	// Create a cookie with the same name as the authentication token cookie.
	// Set its value to an empty string and its expiration time to the past.
	c.Cookie(&fiber.Cookie{
		Name:    "token", // The name of the cookie to be invalidated.
		Value:   "",      // The value is set to an empty string to clear the token.
		Expires: expired, // Set the expiration time to 24 hours ago to ensure the cookie is expired.
	})

	// Return a 200 OK status with a success message.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

/*
To sign out a user, an expired cookie will be sent, effectively deleting the current cookie stored in their API client or web browser.
*/
