package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type GenusController struct {
	GenusService *services.GenusService
}

func NewGenusController(genusService *services.GenusService) *GenusController {
	return &GenusController{
		GenusService: genusService,
	}
}

func (genusController *GenusController) ListGenus(c *fiber.Ctx) error {
	query := &types.ListGenusQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	genus, err := genusController.GenusService.ListGenus(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list genus",
		}
	}
	return c.JSON(&types.ListGenusResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: genus,
	})
}
