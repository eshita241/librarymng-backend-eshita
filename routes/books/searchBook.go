package books

import (
	"librarymng-backend/database"
	"librarymng-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SearchBooks searches for books by title, author, or publisher
func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("query")
	searchType := c.Query("type")

	var books []models.Book
	var result *gorm.DB

	allowedSearchTypes := map[string]bool{
		"name":      true,
		"author":    true,
		"publisher": true,
	}

	if _, ok := allowedSearchTypes[searchType]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid search type",
		})
	}

	result = database.Database.Db.Where(searchType+" LIKE ?", "%"+query+"%").Find(&books)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"books":  books,
	})
}
