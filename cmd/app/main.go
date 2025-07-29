package main

import (
	"log"

	"github.com/Aavishkar-Kolte/go-url-shortner/pkg/global"
	"github.com/Aavishkar-Kolte/go-url-shortner/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	rdb := global.Rdb
	defer rdb.Close()

	app := fiber.New()
	defer app.Shutdown()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/:code", routes.Resolve)
	app.Post("/shorten", routes.Shorten)
}
