package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type StationController struct {
	StationService *services.StationService
}

func NewStationController(stationService *services.StationService) *StationController {
	return &StationController{
		StationService: stationService,
	}
}

func (stationController *StationController) ListStation(c *fiber.Ctx) error {
	query := &types.ListStationQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	station, err := stationController.StationService.ListStation(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list station",
		}
	}
	return c.JSON(&types.ListStationResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: station,
	})
}
