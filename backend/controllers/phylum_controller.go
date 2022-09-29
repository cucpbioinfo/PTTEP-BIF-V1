package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type PhylumController struct {
	PhylumService *services.PhylumService
}

func NewPhylumController(phylumService *services.PhylumService) *PhylumController {
	return &PhylumController{
		PhylumService: phylumService,
	}
}

func (phylumController *PhylumController) ListPhylum(c *fiber.Ctx) error {
	query := &types.ListPhylumQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	phylums, err := phylumController.PhylumService.ListPhylum(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list phylum",
		}
	}
	return c.JSON(&types.ListPhylumResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: phylums,
	})
}
