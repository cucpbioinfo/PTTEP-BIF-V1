package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type ProjectRepository struct {
	pg *pg.DB
}

func NewProjectRepository(pg *pg.DB) *ProjectRepository {
	return &ProjectRepository{
		pg: pg,
	}
}

func (projectRepository *ProjectRepository) ListProject() ([]models.Project, error) {
	var projects []models.Project
	dbQuery := projectRepository.pg.Model(&projects)

	err := dbQuery.Select()
	if err != nil {
		return make([]models.Project, 0), err
	}
	return projects, nil
}
