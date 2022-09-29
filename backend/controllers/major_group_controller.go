package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type MajorGroupController struct {
	MajorGroupService *services.MajorGroupService
}

func NewMajorGroupController(majorGroupService *services.MajorGroupService) *MajorGroupController {
	return &MajorGroupController{
		MajorGroupService: majorGroupService,
	}
}

func (majorGroupController *MajorGroupController) ListMajorGroup(c *fiber.Ctx) error {
	majorGroups, err := majorGroupController.MajorGroupService.ListMajorGroup()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list majorGroup",
		}
	}
	return c.JSON(&types.ListMajorGroupResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: majorGroups,
	})
}
