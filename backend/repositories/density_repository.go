package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/types"
)

type DensityRepository struct {
	pg *pg.DB
}

func NewDensityRepository(pg *pg.DB) *DensityRepository {
	return &DensityRepository{
		pg: pg,
	}
}

func (densityRepository *DensityRepository) ListDensity(query types.ListDensityQuery) ([]models.Density, error) {
	var density []models.Density
	dbQuery := densityRepository.pg.Model(&density)

	if query.AssetID != uuid.Nil {
		dbQuery.Where("density.asset_id = ?", query.AssetID)
	}
	if query.PlatformID != uuid.Nil {
		dbQuery.Where("density.platform_id = ?", query.PlatformID)
	}
	if query.StationID != uuid.Nil {
		dbQuery.Where("density.station_id = ?", query.StationID)
	}
	if query.SpeciesID != uuid.Nil {
		dbQuery.Where("density.species_id = ?", query.SpeciesID)
	}

	err := dbQuery.
		Relation("Asset").
		Relation("Platform").
		Relation("Station").
		Select()
	if err != nil {
		return make([]models.Density, 0), err
	}
	return density, nil
}
