package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type EvennessController struct {
	EvennessService *services.EvennessService
}

func NewEvennessController(evennessService *services.EvennessService) *EvennessController {
	return &EvennessController{
		EvennessService: evennessService,
	}
}

func (evennessController *EvennessController) ListEvenness(c *fiber.Ctx) error {
	evennesses, err := evennessController.EvennessService.ListEvenness()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list evenness",
		}
	}
	return c.JSON(&types.EvennessListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: evennesses,
	})
}
