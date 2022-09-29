package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type MainStationController struct {
	MainStationService *services.MainStationService
}

func NewMainStationController(mainStationService *services.MainStationService) *MainStationController {
	return &MainStationController{
		MainStationService: mainStationService,
	}
}

func (mainStationController *MainStationController) ListMainStation(c *fiber.Ctx) error {
	query := &types.MainStationListQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}

	mainStations, err := mainStationController.MainStationService.ListMainStation(query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	return c.JSON(&types.MainStationListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: mainStations,
	})
}
