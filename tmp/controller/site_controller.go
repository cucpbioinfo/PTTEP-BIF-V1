package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type SiteController struct {
	SiteService *services.SiteService
}

func NewSiteController(siteService *services.SiteService) *SiteController {
	return &SiteController{
		SiteService: siteService,
	}
}

func (siteController *SiteController) ListSite(c *fiber.Ctx) error {
	sites, err := siteController.SiteService.ListSite()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	return c.JSON(&types.SiteListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: sites,
	})
}
