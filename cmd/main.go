package main

import (
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"

	"go-fiber/config"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	app := fiber.New()
	log.SetLevel(log.Level(logConfig.Level))
	app.Use(recover.New())
	home.NewHandler(app) // это вызов пакета home и функции из него, в которую мы ложим app.
	app.Listen(":3000")
}
