package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type EvennessRepository struct {
	pg *pg.DB
}

func NewEvennessRepository(pg *pg.DB) *EvennessRepository {
	return &EvennessRepository{
		pg: pg,
	}
}

func (evennessRepository *EvennessRepository) ListEvenness(query types.ListEvennessQuery) ([]models.Evenness, error) {
	var evenness []models.Evenness
	dbQuery := evennessRepository.pg.Model(&evenness)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("evenness.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("evenness.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("evenness.station_id = ?", query.StationID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Limit(1).
		Select()
	if err != nil {
		return make([]models.Evenness, 0), err
	}
	return evenness, nil
}
