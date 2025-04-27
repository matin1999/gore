package router

import (
	"gore/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, handler handler.HandlerInterface) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Service is up.")
	})

	//order
	app.Post("/order/:pair_id", handler.CreateOrderRequest)
	app.Delete("/order/:id", cancelOrder)
	app.Get("/orderbook", getOrderBook)
	app.Get("/orders/:user_id", getUserOrders)

	//balance
	

}
