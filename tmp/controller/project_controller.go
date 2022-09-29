package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type ProjectController struct {
	ProjectService *services.ProjectService
}

func NewProjectController(projectService *services.ProjectService) *ProjectController {
	return &ProjectController{
		ProjectService: projectService,
	}
}

func (projectController *ProjectController) ListProject(c *fiber.Ctx) error {
	projects, err := projectController.ProjectService.ListProject()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list project",
		}
	}
	return c.JSON(&types.ListProjectResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: projects,
	})
}
