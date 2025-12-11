package main

import (
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"

	"go-fiber/config"

	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	//	app.Get("/", func(c *fiber.Ctx) error {
	//		return c.SendString("Hello")
	//	})
	home.NewHandler(app) // это вызов пакета home и функции из него, в которую мы ложим app.
	app.Listen(":3000")
}
