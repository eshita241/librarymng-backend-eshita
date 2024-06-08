package users

import (
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateUser(c *fiber.Ctx) error {
	var userInput models.Auth
	err := c.BodyParser(&userInput)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return c.Status(400).SendString("Error parsing JSON")
	}

	// if userInput.ID == 0 {
	// 	log.Printf("Invalid user ID provided\n")
	// 	return c.Status(400).SendString("Invalid user ID provided")
	// }

	var updatedUser models.Auth
	resultant := database.Database.Db.First(&updatedUser, userInput.ID)
	if resultant.Error != nil {
		if errors.Is(resultant.Error, gorm.ErrRecordNotFound) {
			log.Printf("User with ID %v not found\n", userInput.ID)
			return c.Status(404).SendString("User not found")
		}
		log.Printf("Error fetching user: %v\n", resultant.Error)
		return c.Status(500).SendString("Error fetching user")
	}

	if err := validator.ValidateUser(userInput); err != nil {
		log.Printf("Error validating user: %v\n", err)
		return c.Status(400).SendString("Error validating user")
	}

	previousPassword := updatedUser.Password

	if userInput.Password != "" && previousPassword != userInput.Password {
		updatedUser.Password = userInput.Password
	} else if previousPassword == userInput.Password {
		return c.Status(400).SendString("you cannot keep the same password:(")
	}

	resultant = database.Database.Db.Save(&updatedUser)
	if resultant.Error != nil {
		log.Printf("Error updating user: %v\n", resultant.Error)
		return c.Status(500).SendString("Error updating user")
	}

	responses := map[string]interface{}{
		"previous_password": previousPassword,
		"updated_user":      updatedUser,
	}

	return c.Status(200).JSON(responses)
}
