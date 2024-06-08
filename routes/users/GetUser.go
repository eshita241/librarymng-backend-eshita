package users

import (
	"errors"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"librarymng-backend/validator"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Printf("error converting user ID to integer: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user models.Auth

	if err := validator.ValidateUser(user); err != nil {
		log.Printf("error validating user: %v\n", err)
		return c.Status(400).SendString("error validating user")
	}

	resultant := database.Database.Db.First(&user, userId)

	if resultant.Error != nil {
		if errors.Is(resultant.Error, gorm.ErrRecordNotFound) {
			log.Printf("user with id %v not found\n", userId)
			return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		}
		log.Printf("error getting user: %v\n", resultant.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	log.Printf("user with ID %v was found\n", userId)
	return c.Status(200).JSON(user)
}
