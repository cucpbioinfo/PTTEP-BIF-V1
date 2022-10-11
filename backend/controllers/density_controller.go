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

// Platform platform
type PlatformDensityController struct {
	PlatformDensityService *services.PlatformDensityService
}

func NewPlatformDensityController(platformdensityService *services.PlatformDensityService) *PlatformDensityController {
	return &PlatformDensityController{
		PlatformDensityService: platformdensityService,
	}
}

func (platformdensityController *PlatformDensityController) ListPlatformDensity(c *fiber.Ctx) error {
	query := &types.ListPlatformDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := platformdensityController.PlatformDensityService.ListPlatformDensity(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListPlatformDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}

// Asset asset
type AssetDensityController struct {
	AssetDensityService *services.AssetDensityService
}

func NewAssetDensityController(assetdensityService *services.AssetDensityService) *AssetDensityController {
	return &AssetDensityController{
		AssetDensityService: assetdensityService,
	}
}

func (assetdensityController *AssetDensityController) ListAssetDensity(c *fiber.Ctx) error {
	query := &types.ListAssetDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := assetdensityController.AssetDensityService.ListAssetDensity(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListAssetDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}
