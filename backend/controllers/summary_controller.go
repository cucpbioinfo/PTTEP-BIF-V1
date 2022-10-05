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
