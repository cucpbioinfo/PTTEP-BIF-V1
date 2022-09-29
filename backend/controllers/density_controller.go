package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type DensityController struct {
	DensityService *services.DensityService
}

func NewDensityController(densityService *services.DensityService) *DensityController {
	return &DensityController{
		DensityService: densityService,
	}
}

func (densityController *DensityController) ListDensity(c *fiber.Ctx) error {
	query := &types.ListDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := densityController.DensityService.ListDensity(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}
