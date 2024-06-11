// SearchBooks searches for books by title, author, or publisher
/*func SearchBooks(c *fiber.Ctx) error {
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
}*/

package books

import (
	"librarymng-backend/database"
	"librarymng-backend/models"

	"github.com/gofiber/fiber/v2"
)

//problem in the above code was that we can have only 2 things, key and value. hence we cannot have what search type as in name or publication or author.
//So the solution is a basic search feature which depends on the word/part of the word.

// SearchBooks handles the book search request
func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter 'q' is required"})
	}

	books, err := searchBooks(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search books"})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// searchBooks queries the database for books matching the search query
/*func searchBooks(query string) ([]models.Book, error) {
	//some error with line rows, err := database.DB.Query("SELECT id, title, author FROM books WHERE title ILIKE $1 OR author ILIKE $1", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		books = append(books, book)
	}
	return books, nil
}*/

// searchBooks queries the database for books matching the search query
func searchBooks(query string) ([]models.Book, error) {
	var books []models.Book
	result := database.Database.Db.Where("title ILIKE ? OR author ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
