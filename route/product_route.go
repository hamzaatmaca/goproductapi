package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/productapi/controller"
)

func ProductRoute(app *fiber.App) {

	app.Get("/product", controller.GetProduct)

	app.Get("/product/:id", controller.GetOneProduct)

	app.Post("/product", controller.AddProduct)

	app.Put("/product/:id", controller.UpdateProduct)

	app.Delete("/product/:id", controller.DeleteProduct)

}
