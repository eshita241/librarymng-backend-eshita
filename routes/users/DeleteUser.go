package users

import (
	"fmt"
	"librarymng-backend/database"
	"librarymng-backend/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		log.Printf("Error converting user id to int: %v\n", err)
		return c.Status(400).SendString("Invalid user id")
	}

	result := database.Database.Db.Delete(&models.AuthResponse{}, userId)

	if result.Error != nil {
		log.Printf("error deleting user: %v\n", result.Error)
		return c.Status(500).SendString("Could not delete user")
	}

	if result.RowsAffected == 0 {
		log.Printf("error deleting user, userid: %v not found\n", userId)
		return c.Status(404).SendString("user not found")
	}

	responseMessage := fmt.Sprintf("User with id: %v deleted", userId)
	log.Println(responseMessage)
	return c.Status(200).SendString(responseMessage)
}
