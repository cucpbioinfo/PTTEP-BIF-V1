package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type DiversityController struct {
	DiversityService *services.DiversityService
}

func NewDiversityController(diversityService *services.DiversityService) *DiversityController {
	return &DiversityController{
		DiversityService: diversityService,
	}
}

func (diversityController *DiversityController) ListDiversity(c *fiber.Ctx) error {
	query := &types.ListDiversityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	diversity, err := diversityController.DiversityService.ListDiversity(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list diversity",
		}
	}
	return c.JSON(&types.ListDiversityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: diversity,
	})
}
