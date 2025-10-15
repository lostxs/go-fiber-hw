package main

import (
	"log"
	"lostx/go-fiber-hw/config"
	"lostx/go-fiber-hw/internal/pages"
	"lostx/go-fiber-hw/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	logger.New(&cfg.Logger)
	app := fiber.New()

	pages.NewHomeHandler(app)

	log.Fatal(app.Listen(":3000"))
}
