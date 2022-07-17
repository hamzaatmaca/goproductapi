package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/productapi/database"
	"github.com/productapi/port"
	"github.com/productapi/route"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
	route.ProductRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is A Rest-API")
	})

	app.Listen(":" + port.PORTLoad())

}
