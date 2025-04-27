package handler

import (
	"fmt"
	"gore/pkg/env"
	"gore/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Envs          *env.Envs
	Metrics       *metrics.Metrics
}




func HandlerInit(envs *env.Envs,metrics *metrics.Metrics) HandlerInterface {
	return &Handler{
		Envs:          envs,
		Metrics:       metrics,
	}
}

func (h *Handler) CreateOrderRequest(c *fiber.Ctx) error {
	fmt.Println("implimentation")
	return fiber.ErrBadGateway
}
