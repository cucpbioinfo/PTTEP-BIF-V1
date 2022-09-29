package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/wichadak/eDNA/models"
)

type PlatformRepository struct {
	pg *pg.DB
}

func NewPlatformRepository(pg *pg.DB) *PlatformRepository {
	return &PlatformRepository{
		pg: pg,
	}
}

func (platformRepository *PlatformRepository) ListPlatform() ([]models.Platform, error) {
	var Platforms []models.Platform
	dbQuery := platformRepository.pg.Model(&Platforms)

	err := dbQuery.
		Relation("Project").
		Select()
	if err != nil {
		return make([]models.Platform, 0), err
	}
	return Platforms, nil
}
