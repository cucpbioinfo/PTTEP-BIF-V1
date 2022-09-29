package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type OccurrenceController struct {
	OccurrenceService *services.OccurrenceService
}

func NewOccurrenceController(occurrenceService *services.OccurrenceService) *OccurrenceController {
	return &OccurrenceController{
		OccurrenceService: occurrenceService,
	}
}

func (occurrenceController *OccurrenceController) ListOccurrence(c *fiber.Ctx) error {
	occurrences, total, err := occurrenceController.OccurrenceService.ListOccurrence()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Bad Request",
		}
	}
	return c.JSON(&types.ListOccurrenceResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data:  occurrences,
		Total: total,
	})
}
