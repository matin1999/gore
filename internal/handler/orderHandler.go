package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)



type OrderHandlerInterface interface {
	CreateOrderRequest(c *fiber.Ctx) error
	CancelOrderRequest(c *fiber.Ctx) error
	OrderBookRequest(c *fiber.Ctx) error
}



func (h *Handler) CreateOrderRequest(c *fiber.Ctx) error {
	fmt.Println("implimentation")
	return fiber.ErrBadGateway
}

func (h *Handler) CreateOrderRequest(c *fiber.Ctx) error {
	fmt.Println("implimentation")
	return fiber.ErrBadGateway
}

func (h *Handler) CreateOrderRequest(c *fiber.Ctx) error {
	fmt.Println("implimentation")
	return fiber.ErrBadGateway
}