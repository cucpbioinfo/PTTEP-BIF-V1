package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type PlatformController struct {
	PlatformService *services.PlatformService
}

func NewPlatformController(platformService *services.PlatformService) *PlatformController {
	return &PlatformController{
		PlatformService: platformService,
	}
}

func (platformController *PlatformController) ListPlatform(c *fiber.Ctx) error {
	query := &types.ListPlatformQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	platform, err := platformController.PlatformService.ListPlatform(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list platform",
		}
	}
	return c.JSON(&types.ListPlatformResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: platform,
	})
}

func (platformController *PlatformController) ListRefPlatform(c *fiber.Ctx) error {
	query := &types.ListPlatformQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	platform, err := platformController.PlatformService.ListRefPlatform(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list platform",
		}
	}
	return c.JSON(&types.ListPlatformResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: platform,
	})
}
