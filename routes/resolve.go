package routes

import (
	"context"
	"log"

	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/global"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func Resolve(c *fiber.Ctx) error {
	rdb := global.Rdb
	shortCode := c.Params("code")
	log.Println("Short code received:", shortCode)

	url, err := rdb.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		log.Fatal("Short code not found in Redis")
		return c.Status(fiber.StatusNotFound).SendString("Short code not found")
	} else if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving URL")
	}

	log.Println("Retrieved URL:", url)
	return c.Redirect(url, fiber.StatusFound)
}
