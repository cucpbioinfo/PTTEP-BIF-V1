package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type ClassController struct {
	ClassService *services.ClassService
}

func NewClassController(classService *services.ClassService) *ClassController {
	return &ClassController{
		ClassService: classService,
	}
}

func (classController *ClassController) ListClass(c *fiber.Ctx) error {
	query := &types.ListClassQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	classes, err := classController.ClassService.ListClass(*query)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list class",
		}
	}
	return c.JSON(&types.ListClassResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: classes,
	})
}
