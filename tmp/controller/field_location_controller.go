package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type FieldLocationController struct {
	FieldLocationService *services.FieldLocationService
}

func NewFieldLocationController(fieldLocationService *services.FieldLocationService) *FieldLocationController {
	return &FieldLocationController{
		FieldLocationService: fieldLocationService,
	}
}

func (fieldLocationController *FieldLocationController) ListFieldLocation(c *fiber.Ctx) error {
	query := &types.FieldLocationListQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}

	fieldLocations, err := fieldLocationController.FieldLocationService.ListFieldLocation(query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list fieldLocation",
		}
	}
	return c.JSON(&types.FieldLocationListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: fieldLocations,
	})
}
