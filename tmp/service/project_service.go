package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type ProjectService struct {
	ProjectRepository *repositories.ProjectRepository
}

func NewProjectService(projectRepository *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepository: projectRepository,
	}
}

func (projectService *ProjectService) ListProject() ([]types.ProjectDto, error) {
	projects, err := projectService.ProjectRepository.ListProject()
	if err != nil {
		return make([]types.ProjectDto, 0), err
	}
	projectList := make([]types.ProjectDto, len(projects))
	for i, project := range projects {
		projectList[i] = types.ProjectDto{
			ProjectID:   project.ProjectID,
			ProjectName: project.ProjectName,
		}
	}
	return projectList, nil
}
