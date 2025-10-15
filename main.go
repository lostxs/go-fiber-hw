package main

import (
	"log"
	"lostx/go-fiber-hw/config"
	"lostx/go-fiber-hw/internal/pages"
	"lostx/go-fiber-hw/pkg/logger"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	cfg := config.Load()
	logger := logger.New(&cfg.Logger)
	app := fiber.New()

	app.Use(slogfiber.New(logger))

	pages.NewHomeHandler(app)

	log.Fatal(app.Listen(":3000"))
}
