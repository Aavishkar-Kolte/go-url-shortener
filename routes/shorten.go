package routes

import (
	"context"
	"log"
	"time"

	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/global"
	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type shortenRequest struct {
	Url string `form:"url"`
}

type shortenResponse struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	ExpiresAt   time.Time `json:"expires_at" validate:"required"`
}

func Shorten(c *fiber.Ctx) error {
	rdb := global.Rdb
	var req shortenRequest

	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	log.Println("/shorten", req)

	if req.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("URL is required")
	}

	shortURL, err := utils.RandomBase62(6)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate short URL")
	}

	err = rdb.Set(context.Background(), shortURL, req.Url, 180 * 24 * time.Hour).Err()
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to shorten URL")
	}

	log.Println("Shortened URL:", shortURL)

	return c.JSON(shortenResponse{
		ShortURL:    shortURL,
		OriginalURL: req.Url,
		ExpiresAt:   time.Now().Add(180 * 24 * time.Hour),
	})
}
