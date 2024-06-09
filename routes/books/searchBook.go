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

	switch searchType {
	case "title":
		result = database.Database.Db.Where("title LIKE ?", "%"+query+"%").Find(&books)
	case "author":
		result = database.Database.Db.Where("author LIKE ?", "%"+query+"%").Find(&books)
	case "publisher":
		result = database.Database.Db.Where("publisher LIKE ?", "%"+query+"%").Find(&books)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid search type",
		})
	}

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
