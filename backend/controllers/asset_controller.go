package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type AssetController struct {
	AssetService *services.AssetService
}

func NewAssetController(assetService *services.AssetService) *AssetController {
	return &AssetController{
		AssetService: assetService,
	}
}

func (assetController *AssetController) ListAsset(c *fiber.Ctx) error {
	assets, err := assetController.AssetService.ListAsset()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list asset",
		}
	}
	return c.JSON(&types.ListAssetResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: assets,
	})
}
