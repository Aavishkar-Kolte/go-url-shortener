package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	app := fiber.New()
	defer app.Shutdown()

	app.Get("/:code", func(c *fiber.Ctx) error {
		shortCode := c.Params("code")
		log.Println("Short code received:", shortCode)

		url, err := rdb.Get(ctx, shortCode).Result()
		if err == redis.Nil {
			log.Fatal("Short code not found in Redis")
			return c.Status(fiber.StatusNotFound).SendString("Short code not found")
		} else if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving URL")
		}

		log.Println("Retrieved URL:", url)
		return c.Redirect(url, fiber.StatusFound)
	})

	app.Post("/shorten", func(c *fiber.Ctx) error {
		log.Println("/shorten", c.FormValue("url"))
		url := c.FormValue("url")
		if url == "" {
			return c.Status(fiber.StatusBadRequest).SendString("URL is required")
		}

		shortURL := generateShortURL()

		err := rdb.Set(ctx, shortURL, url, 0).Err()
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to shorten URL")
		}

		log.Println("Shortened URL:", shortURL)

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}

func generateShortURL() string {
	// Placeholder function to generate a short URL
	return uuid.New().String()
}
