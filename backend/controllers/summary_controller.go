package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type SummaryController struct {
	SummaryService *services.SummaryService
}

func NewSummaryController(summaryService *services.SummaryService) *SummaryController {
	return &SummaryController{
		SummaryService: summaryService,
	}
}

func (summaryController *SummaryController) ListSummary(c *fiber.Ctx) error {
	query := &types.ListSummaryQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	summary, err := summaryController.SummaryService.ListSummary(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list summary",
		}
	}
	return c.JSON(&types.ListSummaryResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: summary,
	})
}

// all summary
func (summaryController *SummaryController) ListAllSummary(c *fiber.Ctx) error {
	query := &types.ListSummaryQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	summary, err := summaryController.SummaryService.ListAllSummary(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list summary",
		}
	}
	return c.JSON(&types.ListSummaryResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: summary,
	})
}

// Platform platform
type PlatformSummaryController struct {
	PlatformSummaryService *services.PlatformSummaryService
}

func NewPlatformSummaryController(platformsummaryService *services.PlatformSummaryService) *PlatformSummaryController {
	return &PlatformSummaryController{
		PlatformSummaryService: platformsummaryService,
	}
}

func (platformsummaryController *PlatformSummaryController) ListPlatformSummary(c *fiber.Ctx) error {
	query := &types.ListPlatformSummaryQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	summary, err := platformsummaryController.PlatformSummaryService.ListPlatformSummary(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list platformsummary",
		}
	}
	return c.JSON(&types.ListPlatformSummaryResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: summary,
	})
}

// Asset asset
type AssetSummaryController struct {
	AssetSummaryService *services.AssetSummaryService
}

func NewAssetSummaryController(assetsummaryService *services.AssetSummaryService) *AssetSummaryController {
	return &AssetSummaryController{
		AssetSummaryService: assetsummaryService,
	}
}

func (assetsummaryController *AssetSummaryController) ListAssetSummary(c *fiber.Ctx) error {
	query := &types.ListAssetSummaryQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	summary, err := assetsummaryController.AssetSummaryService.ListAssetSummary(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list assetsummary",
		}
	}
	return c.JSON(&types.ListAssetSummaryResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: summary,
	})
}

// Year year
type YearSummaryController struct {
	YearSummaryService *services.YearSummaryService
}

func NewYearSummaryController(yearsummaryService *services.YearSummaryService) *YearSummaryController {
	return &YearSummaryController{
		YearSummaryService: yearsummaryService,
	}
}

func (yearsummaryController *YearSummaryController) ListYearSummary(c *fiber.Ctx) error {
	query := &types.ListYearSummaryQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	yearsummary, err := yearsummaryController.YearSummaryService.ListYearSummary(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list yearsummary",
		}
	}
	return c.JSON(&types.ListYearSummaryResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: yearsummary,
	})
}
