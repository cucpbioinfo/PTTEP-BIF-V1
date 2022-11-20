package controllers

import (
	"fmt"

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

// alldensity
func (densityController *DensityController) ListAllDensity(c *fiber.Ctx) error {
	query := &types.ListDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := densityController.DensityService.ListAllDensity(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list alldensity",
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
		fmt.Println("er", err)
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

// Year Asset year asset
type YearAssetDensityController struct {
	YearAssetDensityService *services.YearAssetDensityService
}

func NewYearAssetDensityController(yearassetdensityService *services.YearAssetDensityService) *YearAssetDensityController {
	return &YearAssetDensityController{
		YearAssetDensityService: yearassetdensityService,
	}
}

func (yearassetdensityController *YearAssetDensityController) ListYearAssetDensity(c *fiber.Ctx) error {
	query := &types.ListYearAssetDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := yearassetdensityController.YearAssetDensityService.ListYearAssetDensity(*query)
	if err != nil {
		fmt.Println("er", err)
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListYearAssetDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}

// Year Platform year platform
type YearPlatformDensityController struct {
	YearPlatformDensityService *services.YearPlatformDensityService
}

func NewYearPlatformDensityController(yearplatformdensityService *services.YearPlatformDensityService) *YearPlatformDensityController {
	return &YearPlatformDensityController{
		YearPlatformDensityService: yearplatformdensityService,
	}
}

func (yearplatformdensityController *YearPlatformDensityController) ListYearPlatformDensity(c *fiber.Ctx) error {
	query := &types.ListYearPlatformDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := yearplatformdensityController.YearPlatformDensityService.ListYearPlatformDensity(*query)
	if err != nil {
		fmt.Println("er", err)
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListYearPlatformDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}

// Year Station year stationn
type YearDensityController struct {
	YearDensityService *services.YearDensityService
}

func NewYearDensityController(yeardensityService *services.YearDensityService) *YearDensityController {
	return &YearDensityController{
		YearDensityService: yeardensityService,
	}
}

func (yeardensityController *YearDensityController) ListYearDensity(c *fiber.Ctx) error {
	query := &types.ListYearDensityQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	density, err := yeardensityController.YearDensityService.ListYearDensity(*query)
	if err != nil {
		fmt.Println("er", err)
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list density",
		}
	}
	return c.JSON(&types.ListYearDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: density,
	})
}
