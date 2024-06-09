package middleware

import (
	"encoding/json"
	"librarymng-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func CacheMiddleware(cacheInstance *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			// This makes only GET requests go through cache middleware, for other methods it will follow the next middleware
			return c.Next()
		}

		cacheKey := c.Path() + "?" + c.Params("id")

		if cached, found := cacheInstance.Get(cacheKey); found {
			return c.JSON(cached)
		}

		err := c.Next()
		if err != nil {
			return err
		}

		body := c.Response().Body()

		var book models.Book
		err = json.Unmarshal(body, &book)
		if err != nil {
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		// Cache the response for a limited time
		cacheInstance.Set(cacheKey, book, 10*time.Minute)

		return nil
	}
}
