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
	query := &types.DiversityListQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	diversities, err := diversityController.DiversityService.ListDiversity(types.DiversityListQuery{
		Year:       query.Year,
		ProjectID:  query.ProjectID,
		PlatformID: query.PlatformID,
		StationID:  query.StationID,
	})
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list species",
		}
	}
	return c.JSON(&types.DiversityListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: diversities,
	})
}
