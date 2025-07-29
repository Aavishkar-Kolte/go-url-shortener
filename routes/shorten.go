package routes

import (
	"context"
	"log"

	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/global"
	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Shorten(c *fiber.Ctx) error {
	rdb := global.Rdb
	log.Println("/shorten", c.FormValue("url"))
	url := c.FormValue("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("URL is required")
	}

	shortURL, err := utils.RandomBase62(6)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate short URL")
	}

	err = rdb.Set(context.Background(), shortURL, url, 0).Err()
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to shorten URL")
	}

	log.Println("Shortened URL:", shortURL)

	return nil
}
