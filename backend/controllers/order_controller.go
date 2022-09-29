package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type OrderController struct {
	OrderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

func (orderController *OrderController) ListOrder(c *fiber.Ctx) error {
	query := &types.ListOrderQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	orders, err := orderController.OrderService.ListOrder(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list order",
		}
	}
	return c.JSON(&types.ListOrderResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: orders,
	})
}
