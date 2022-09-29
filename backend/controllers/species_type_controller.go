package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/types"
)

type SpeciesTypeController struct{}

func NewSpeciesTypeController() *SpeciesTypeController {
	return &SpeciesTypeController{}
}

func (speciesTypeController *SpeciesTypeController) ListSpeciesType(c *fiber.Ctx) error {
	return c.JSON(&types.SpeciesTypeListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: types.SpeciesTypes[:],
	})
}
