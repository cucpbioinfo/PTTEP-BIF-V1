package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type DiversityRepository struct {
	pg *pg.DB
}

func NewDiversityRepository(pg *pg.DB) *DiversityRepository {
	return &DiversityRepository{
		pg: pg,
	}
}
func (DiversityRepository *DiversityRepository) ListDiversity(query types.DiversityListQuery) ([]models.Diversity, error) {
	var Diversities []models.Diversity
	dbquery := DiversityRepository.pg.Model(&Diversities)
	if query.ProjectID != uuid.Nil {
		dbquery.Where("diversity.project_id = ?", query.ProjectID)
	}
	if query.PlatformID != uuid.Nil {
		dbquery.Where("diversity.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbquery.Where("diversity.station_id = ?", query.StationID)
	}
	err := dbquery.
		Relation("Project").
		Relation("Platform").
		Relation("Station").
		Select()
	if err != nil {
		return make([]models.Diversity, 0), err
	}
	return Diversities, nil
}
