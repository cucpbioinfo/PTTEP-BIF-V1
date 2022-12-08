package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type LocationController struct {
	LocationService *services.LocationService
}

func NewLocationController(locationService *services.LocationService) *LocationController {
	return &LocationController{
		LocationService: locationService,
	}
}

func (locationController *LocationController) ListLocation(c *fiber.Ctx) error {
	query := &types.ListLocationQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	location, err := locationController.LocationService.ListLocation(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list location",
		}
	}
	return c.JSON(&types.ListLocationResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: location,
	})
}

func (locationController *LocationController) ListLocationAsset(c *fiber.Ctx) error {
	query := &types.ListLocationQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	location, err := locationController.LocationService.ListLocationAsset(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list location",
		}
	}
	return c.JSON(&types.ListLocationAssetResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: location,
	})
}

func (locationController *LocationController) ListLocationPlatform(c *fiber.Ctx) error {
	query := &types.ListLocationQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	location, err := locationController.LocationService.ListLocationPlatform(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list location",
		}
	}
	return c.JSON(&types.ListLocationPlatformResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: location,
	})
}
