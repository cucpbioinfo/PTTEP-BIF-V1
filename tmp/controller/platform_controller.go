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
	platforms, err := platformController.PlatformService.ListPlatform()
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
		Data: platforms,
	})
}
