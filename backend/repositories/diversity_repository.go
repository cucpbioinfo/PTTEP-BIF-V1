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

func (diversityRepository *DiversityRepository) ListDiversity(query types.ListDiversityQuery) ([]models.Diversity, error) {
	var diversity []models.Diversity
	dbQuery := diversityRepository.pg.Model(&diversity)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("diversity.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("diversity.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("diversity.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.Diversity, 0), err
	}
	return diversity, nil
}
