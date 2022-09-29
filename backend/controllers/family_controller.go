package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type FamilyController struct {
	FamilyService *services.FamilyService
}

func NewFamilyController(familyService *services.FamilyService) *FamilyController {
	return &FamilyController{
		FamilyService: familyService,
	}
}

func (familyController *FamilyController) ListFamily(c *fiber.Ctx) error {
	query := &types.ListFamilyQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	families, err := familyController.FamilyService.ListFamily(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list family",
		}
	}
	return c.JSON(&types.ListFamilyResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: families,
	})
}
