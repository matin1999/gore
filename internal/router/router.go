package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx,handler handler.HandlerInterface) error {
		return c.JSON("Service is up.")
	})


	app.Post("/order", placeOrder)
    app.Delete("/order/:id", cancelOrder)
    app.Get("/orderbook", getOrderBook)
    app.Get("/orders/:user_id", getUserOrders)

	

}
